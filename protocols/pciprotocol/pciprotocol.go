package pciprotocol

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

type PCIProtocol struct {
	SerialPort         string
	DataMessageReceive chan DataMessage
	AppInitFunction    func(uint8)
	id                 uint8
	port               serial.Port
	rxBuffer           []byte
	rxChannel          chan PCIPacket
	txChannel          chan PCIPacket
	state              safePciState
	pciChannels        safePciChannels
	testRingTimeout    safeTime
	serviceStop        chan bool
}

// NOTE: rxBuffer must only be dealt with by 1 go-routine

func (p *PCIProtocol) Start() {
	// Initialize state
	p.state.Set(PCI_START)

	// Start services
	go p.protocolStateService(p.serviceStop)
	go p.protocolService(p.serviceStop)
	go p.rxService(p.serviceStop)
	go p.txService(p.serviceStop)
}

func (p *PCIProtocol) Stop() {
	p.serviceStop <- true

	p.state.Set(PCI_START)
}

func (p *PCIProtocol) Restart() {
	if p.state.Get() != PCI_START {
		p.Stop()
	}

	p.Start()
}

func (p *PCIProtocol) SendDataMessage(group byte, params []byte) {
	controllerID := p.pciChannels.GetController(group)

	if controllerID <= 0 {
		return
	}

	txParamsLength := len(params) + 1
	txParams := make([]byte, txParamsLength)
	txParams[0] = group & 0x07
	for i := 1; i < txParamsLength; i++ {
		txParams[i] = params[i-1]
	}

	p.txChannel <- PCIPacket{
		source_ID:      p.id,
		destination_ID: uint8(controllerID),
		command:        DATA_MESSAGE,
		parameters:     txParams,
	}
}

// go-routine which handles the PCI connection service
func (p *PCIProtocol) protocolStateService(stop chan bool) {
	for {
		// Handle stop
		select {
		case <-stop:
			return
		default:
			// Check modes
			switch p.state.Get() {
			case PCI_START:
				// Check if serial port init has run
				if p.port == nil {
					// Enforce external settings being configured correctly
					goodConfig := false
					portList, _ := serial.GetPortsList()
					for _, val := range portList {
						if val == p.SerialPort {
							goodConfig = true
						}
					}
					if !goodConfig {
						return
					}
					// Configure serial port
					serialMode := serial.Mode{
						BaudRate: 9600,
						DataBits: 8,
						Parity:   serial.NoParity,
						StopBits: serial.OneStopBit,
					}
					port, err := serial.Open(p.SerialPort, &serialMode)
					if err != nil {
						log.Println(err)
					}
					p.port = port
				}

				// Set PCI ID
				if p.id == 0 {
					p.id = 1
				}

				// Configure channels
				if p.rxChannel == nil {
					p.rxChannel = make(chan PCIPacket, 100)
				}
				if p.txChannel == nil {
					p.txChannel = make(chan PCIPacket, 100)
				}

				p.state.Set(PCI_WAIT_TEST_RING)
			case PCI_WAIT_TEST_RING:

			case PCI_WAIT_ASSIGN:
				if p.testRingTimeout.GetTimeSince().Seconds() > 2.5 {
					// Reset state
					p.pciChannels.ResetChannels()
					p.state.Set(PCI_WAIT_TEST_RING)
				}
			case PCI_CONNECTED:
				if p.testRingTimeout.GetTimeSince().Seconds() > 2.5 {
					// Reset state
					p.pciChannels.ResetChannels()
					p.state.Set(PCI_WAIT_TEST_RING)
				}
			}
		}
	}
}

// go-routine to handle processing packets
func (p *PCIProtocol) protocolService(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			if p.state.Get() != PCI_START {
				for pkt := range p.rxChannel {
					p.packetArrived(pkt)
				}
			}
		}
	}
}

// go-routine which handles decoding incoming messages
func (p *PCIProtocol) rxService(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			if p.state.Get() != PCI_START {
				p.updateReceiveBuffer()
			}
		}
	}
}

func (p *PCIProtocol) txService(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			if p.state.Get() != PCI_START {
				for pkt := range p.txChannel {
					p.sendPacket(pkt)
				}
			}
		}
	}
}

func (p *PCIProtocol) packetArrived(pkt PCIPacket) {
	fmt.Println("Packet arrived:", pkt)
	switch pkt.command {
	case DATA_MESSAGE:
		p.handleDataMessage(pkt)
	case TEST_RING:
		p.handleTestRing(pkt)
	case ASSIGN_TO_GROUP:
		p.handleAssignToGroup(pkt)
	case CONNECTED_TO_PART:
		// Pass on message
		p.txChannel <- pkt
	case DEASSIGN_FROM_GROUP:
		p.handleDeassignFromGroup(pkt)
	case MULTICAST_MODE:
		// pass on message
		p.txChannel <- pkt
	case MULTIPLEXED:
		// Do not implement
	}
}

// Function which handles receiving serial packets
func (p *PCIProtocol) updateReceiveBuffer() {
	// Move data to temporary buffer
	tempBuff := make([]byte, 100)
	len, _ := p.port.Read(tempBuff)

	// Move items to overall receive buffer
	p.rxBuffer = append(p.rxBuffer, tempBuff[:len]...)

	// Check and update receive buffer
	headerIndx := -1

	for i, val := range p.rxBuffer {
		// Check if header byte
		if val&0x80 != 0 {
			headerIndx = i
		}

		// Check if terminator byte
		if val == 13 && headerIndx >= 0 {
			// Remove header bit
			p.rxBuffer[headerIndx] &= 0x7F
			// Extract message
			uuEncodedMessage := p.rxBuffer[headerIndx:i]
			// trim buffer
			p.rxBuffer = p.rxBuffer[i+1:]
			// Decode message
			decodedMsg := uudecode(uuEncodedMessage)
			//fmt.Printf("Received decoded message: [% x]\n", decodedMsg)
			// Convert to PCI Packet
			pkt, err := messageToPacket(decodedMsg)
			if err != nil {
				log.Printf("Malformed packet received. Encoded %[1]v, Decoded: %[2]v", uuEncodedMessage, decodedMsg)
			}
			// Store arrived packet
			p.rxChannel <- pkt
		}
	}
}

func (p *PCIProtocol) sendPacket(pkt PCIPacket) {
	fmt.Println("Sending packet:", pkt)
	// Get encoded message
	dec_msg := pkt.packetToMessage()
	enc_msg := uuencode(dec_msg)
	//fmt.Printf("Sending decoded message: [% x]\n", dec_msg)

	// Add header sync bit
	enc_msg[0] |= 0x80
	// Add termination byte
	enc_msg = append(enc_msg, byte(13))

	// Transmit packet
	_, err := p.port.Write(enc_msg)

	if err != nil {
		log.Println("Error writing packet:", err)
	}
}

func (p *PCIProtocol) handleDataMessage(pkt PCIPacket) {
	// Get parameters
	group := pkt.parameters[0] & 0x07

	// Check if intended recepient
	if pkt.destination_ID != p.id && pkt.destination_ID != 0xF {
		p.txChannel <- pkt
		return
	}

	controllerID := p.pciChannels.GetChannelID(group)

	if controllerID <= 0 || controllerID != int8(pkt.source_ID) {
		p.txChannel <- pkt
		return
	}

	p.DataMessageReceive <- DataMessage{
		Group:  group,
		Params: pkt.parameters,
	}
}

func (p *PCIProtocol) handleTestRing(pkt PCIPacket) {
	// Reset timeout
	p.testRingTimeout.SetNow()

	// Get parameters
	flag := pkt.parameters[0]
	incr := pkt.parameters[1]

	resetRequestFlag := flag&0x01 != 0
	moduleTypePCFlag := flag&0x02 != 0
	moduleTypePCIFlag := flag&0x04 != 0
	resetRingFlag := flag&0x80 != 0

	// Handle reset request from controller
	if resetRingFlag {
		p.pciChannels.ResetChannels()
		p.state.Set(PCI_WAIT_ASSIGN)
	}

	// Check if this is the first test ring
	if p.state.Get() == PCI_WAIT_TEST_RING {
		// Request a reset
		resetRequestFlag = true
	}

	// Check if id match
	if (pkt.destination_ID == p.id) || (pkt.destination_ID == 0xF) {
		// Inform controller this is a camera / PCI module
		moduleTypePCIFlag = true
	}

	// Check for handling incr
	if incr == 0 && ((pkt.destination_ID == p.id) || (pkt.destination_ID == 0xF)) {
		incr++
	} else if incr != 0 {
		incr++
	}

	// Create new flags
	txFlag := byte(0)

	if resetRequestFlag {
		txFlag |= 0x01
	}
	if moduleTypePCFlag {
		txFlag |= 0x02
	}
	if moduleTypePCIFlag {
		txFlag |= 0x04
	}
	if resetRingFlag {
		txFlag |= 0x80
	}

	// Pass new packet
	p.txChannel <- PCIPacket{
		source_ID:      pkt.source_ID,
		destination_ID: pkt.destination_ID,
		command:        TEST_RING,
		parameters: []byte{
			txFlag,
			incr,
		},
	}
}

func (p *PCIProtocol) handleAssignToGroup(pkt PCIPacket) {
	group := pkt.parameters[0] & 0x0F
	//noInitRequest := pkt.parameters[0]&0x80 != 0

	if pkt.destination_ID != p.id {
		// pass packet along
		p.txChannel <- pkt
		return
	}

	chanID := p.pciChannels.NewGroup(group, pkt.source_ID)

	if chanID < 0 {
		// pass packet along, no space in channels
		p.txChannel <- pkt
		return
	}

	// Inform controller with CONNECTED_TO_PART
	stat := byte(0x00)                // Channel #
	stat |= (byte(chanID) & 0x3) << 4 // Is a camera

	p.txChannel <- PCIPacket{
		source_ID:      p.id,
		destination_ID: pkt.source_ID,
		command:        CONNECTED_TO_PART,
		parameters: []byte{
			group,
			stat,
		},
	}

	// Check and update state
	if p.state.Get() != PCI_CONNECTED {
		p.AppInitFunction(group)
		p.state.Set(PCI_CONNECTED)
	}
}

func (p *PCIProtocol) handleDeassignFromGroup(pkt PCIPacket) {
	group := pkt.parameters[0]

	// Check if D_ID matches
	if pkt.destination_ID != p.id {
		// pass on message
		p.txChannel <- pkt
		return
	}

	chanID := p.pciChannels.GetChannelID(group)
	controller := p.pciChannels.GetController(group)

	if controller <= 0 || controller != int8(pkt.source_ID) {
		// pass on message
		p.txChannel <- pkt
		return
	}

	// Deassign from channel
	p.pciChannels.DeleteGroup(group)

	// Inform controller with CONNECTED_TO_PART
	stat := byte(chanID) & 0x3 << 4 // Channel #
	stat &= 0xFD                    // camera not present

	p.txChannel <- PCIPacket{
		source_ID:      p.id,
		destination_ID: pkt.source_ID,
		command:        CONNECTED_TO_PART,
		parameters: []byte{
			group,
			stat,
		},
	}

	// Update state
	if p.state.Get() == PCI_CONNECTED {
		p.state.Set(PCI_WAIT_ASSIGN)
	}
}
