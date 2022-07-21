package pci

import (
	"log"

	"go.bug.st/serial"
)

// Enum for PCI command types
type PCI_CMD byte

const (
	DATA_MESSAGE        PCI_CMD = 0
	TEST_RING           PCI_CMD = 1
	ASSIGN_TO_GROUP     PCI_CMD = 2
	CONNECTED_TO_PART   PCI_CMD = 3
	DEASSIGN_FROM_GROUP PCI_CMD = 4
	MULTICAST_MODE      PCI_CMD = 5
	MULTIPLEXED         PCI_CMD = 7
)

// Enum for PCI connection state
type PCI_CONNECTION_STATE int

const (
	PCI_AWAIT_START    PCI_CONNECTION_STATE = 0
	PCI_WAIT_TEST_RING PCI_CONNECTION_STATE = 1
	PCI_WAIT_ASSIGN    PCI_CONNECTION_STATE = 2
	PCI_CONNECTED      PCI_CONNECTION_STATE = 3
)

// General enums
const (
	GROUP_NONE byte = 0x10
)

// Struct for channels
type Channel struct {
	group byte
	pc_id byte
}

type PCI struct {
	port                  serial.Port
	rxBuffer              []byte
	dataMessageHandler    func(byte, byte, []byte)
	initConnectionHandler func(byte, byte)
	testRingInitialized   bool
	id                    byte
	channels              [4]Channel
	connectionState       PCI_CONNECTION_STATE
}

func (p *PCI) SetPort(name string, destinationID byte) {
	// Configure serial port for PCI
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	// Open port
	port, err := serial.Open(name, mode)
	if err != nil {
		// Failure
		log.Fatal(err)
	}
	// Handle opened port
	p.port = port
	p.testRingInitialized = false
	p.id = destinationID

	for i := 0; i < len(p.channels); i++ {
		p.channels[i] = Channel{
			group: GROUP_NONE,
			pc_id: 0,
		}
	}

	// Set connection state
	p.connectionState = PCI_AWAIT_START
}

// Function to link a function as a datamessage handler
func (p *PCI) SetDataMessageHandler(function func(byte, byte, []byte)) {
	p.dataMessageHandler = function
}

// Function to link a function as an init handler
func (p *PCI) SetInitConnectionHandler(function func(byte, byte)) {
	p.initConnectionHandler = function
}

// Function to send a data message
func (p *PCI) SendDataMessage(d_id byte, group byte, command byte, params []byte) {
	var txParams []byte
	txParams = append(txParams, group&0x07)
	txParams = append(txParams, command)
	txParams = append(txParams, params...)

	var pkt PCIPacket
	pkt.command = DATA_MESSAGE
	pkt.source_ID = p.id
	pkt.destination_ID = d_id
	pkt.parameters = txParams

	//fmt.Println("Sending data message: ", pkt)

	p.sendPacket(pkt)
}

// Function to send a PCIPacket
func (p *PCI) sendPacket(pkt PCIPacket) (err error) {
	// Encode into byte slice
	//fmt.Printf("Sending packet: %[1]v\n", pkt)
	dec_msg := pkt.packetToMessage()
	//fmt.Printf("Sending decoded message: [% x]\n", dec_msg)
	enc_msg := uuencode(dec_msg)

	// Add header sync bit
	enc_msg[0] |= 0x80
	// Add termination byte
	enc_msg = append(enc_msg, 13)

	// Transmit packet
	_, err = p.port.Write(enc_msg)

	return
}

// Handles the opened port, receiving messages, and handling those messages
// Should be called periodically
func (p *PCI) HandleData() {
	// Check and update connection state
	if p.connectionState == PCI_AWAIT_START {
		p.connectionState = PCI_WAIT_TEST_RING
	}

	// Read to temporary buffer
	tempBuff := make([]byte, 100)
	len, _ := p.port.Read(tempBuff)

	// Move items to receive buffer
	for i := 0; i < len; i++ {
		p.rxBuffer = append(p.rxBuffer, tempBuff[i])
	}

	// Check for packets in receive buffer
	headerIndx := 0
	for i, item := range p.rxBuffer {
		// Check if header
		if item&0x80 != 0 {
			headerIndx = i
			// Remove header bit
			p.rxBuffer[i] &= 0x7F
		}
		// Check if terminator
		if item == 13 {
			// Extract message
			enc_msg := msgExtract(p.rxBuffer, headerIndx, i)
			// Remove message from buffer
			p.rxBuffer = append(p.rxBuffer[:headerIndx], p.rxBuffer[i+1:]...)
			// Decode message
			dec_msg := uudecode(enc_msg)
			//fmt.Printf("Received decoded message: [% x]\n", dec_msg)
			// Convert to PCI Packet
			pkt, err := messageToPacket(dec_msg)
			if err != nil {
				log.Printf("Malformed packet received. Encoded: %[1]v, Decoded: %[2]v", enc_msg, dec_msg)
			}
			// Announce that packet arrived
			p.packetArrived(pkt)
			// Reset loop since we messed with the buffer
			break
		}

	}
}

// Function to extract the message from a larger buffer
func msgExtract(src []byte, start int, end int) (dst []byte) {
	for i := start; i < end; i++ {
		dst = append(dst, src[i])
	}

	return
}

// Function to handle when a packet arrives
// Will first handle PCI-protocol specified tasks, then offload to handler
func (p *PCI) packetArrived(pkt PCIPacket) {
	//fmt.Printf("Packet Arrived: %[1]v\n", pkt)

	switch pkt.command {
	case DATA_MESSAGE:
		p.handleDataMessage(pkt)
	case TEST_RING:
		p.handleTestRing(pkt)
	case ASSIGN_TO_GROUP:
		p.handleAssignToGroup(pkt)
	case CONNECTED_TO_PART:
		p.handleConnectedToPart(pkt)
	case DEASSIGN_FROM_GROUP:
		p.handleDeassignFromGroup(pkt)
	case MULTICAST_MODE:

	case MULTIPLEXED:

	}
}

// Function to handle a data message being received
func (p *PCI) handleDataMessage(pkt PCIPacket) {
	// Extract info
	group := pkt.parameters[0] & 0x07
	//msg_baseStation := (pkt.parameters[0] & 0x08) >> 3
	//msg_camera := (pkt.parameters[0] & 0x10) >> 4
	//noAck := (pkt.parameters[0] & 0x80) >> 7

	// Ensure intended recepient
	if p.intendedRecepient(pkt) {
		chanIndex := p.getChannelFromGroup(group)
		if chanIndex >= 0 {
			if p.channels[chanIndex].pc_id == pkt.source_ID {
				// Intendend Recepient
				params := pkt.parameters[1:]
				p.dataMessageHandler(pkt.source_ID, group, params)
				return
			}
		}
	}

	// Not intended recepient, retransmit
	p.sendPacket(pkt)

}

// Function to handle a test ring being received
func (p *PCI) handleTestRing(pkt PCIPacket) {
	// Get parameters
	flag := pkt.parameters[0]
	incr := pkt.parameters[1]

	// Get bools from flag
	//resetRequest := (flag&0x01 != 0)
	//moduleTypePC := (flag&0x02 != 0)
	//moduleTypePCI := (flag&0x04 != 0)
	resetRing := (flag&0x80 != 0)

	// Handle reset request
	if resetRing {
		// Reset all channel assignments
		p.testRingInitialized = true
	}

	// Ensure ring has already been reset
	if !p.testRingInitialized {
		// Set reset request bit
		flag |= 0x01
	}

	// Set PCI type bit if d_id match
	if p.intendedRecepient(pkt) {
		// Set MODULE_TYPE_PCI bit
		flag |= 0x04
	}

	// Handle incr
	if incr == 0 {
		if p.intendedRecepient(pkt) {
			// Increment incr
			incr++
		}
	} else {
		// Increment incr
		incr++
	}

	// Re-transmit same / modified packet
	var txParams []byte
	txParams = append(txParams, flag, incr)

	txPkt := PCIPacket{
		source_ID:      pkt.source_ID,
		destination_ID: pkt.destination_ID,
		command:        TEST_RING,
		parameters:     txParams,
	}

	p.sendPacket(txPkt)

	// Check and update connection state
	if p.connectionState == PCI_WAIT_TEST_RING {
		p.connectionState = PCI_WAIT_ASSIGN
	}
}

// Handles ASSIGN_TO_GROUP messages
func (p *PCI) handleAssignToGroup(pkt PCIPacket) {
	grpByte := pkt.parameters[0]
	grpNum := grpByte & 0x0F
	//noInit := ((grpByte & 0x80) != 0)

	chanID := p.getFreeChannel()

	if !p.intendedRecepient(pkt) || chanID < 0 {
		// Cannot execute command, pass along instead
		p.sendPacket(pkt)
		return
	}

	// Connect to channel
	p.channels[chanID].group = grpNum
	p.channels[chanID].pc_id = pkt.source_ID

	stat := byte(0x00) // Channel # (5-4), Camera (1), Base station (0)
	stat |= byte(chanID) & 0x3 << 4

	// Inform PC device
	var connectedParams []byte
	connectedParams = append(connectedParams, grpNum, stat)

	connectedPkt := PCIPacket{
		source_ID:      p.id,
		destination_ID: pkt.source_ID,
		command:        CONNECTED_TO_PART,
		parameters:     connectedParams,
	}

	p.sendPacket(connectedPkt)

	// Check and update connection state
	if p.connectionState == PCI_WAIT_ASSIGN {
		p.connectionState = PCI_CONNECTED
		// Inform caller
		p.initConnectionHandler(pkt.source_ID, grpNum)
	}
}

// Handle CONNECTED_TO_PART messages
func (p *PCI) handleConnectedToPart(pkt PCIPacket) {
	if !p.intendedRecepient(pkt) {
		// Rebroadcast
		p.sendPacket(pkt)
	}
}

// Handle DEASSIGN_FROM_GROUP messages
func (p *PCI) handleDeassignFromGroup(pkt PCIPacket) {
	//grp := pkt.parameters[0]
}

// Encodes a message and returns the encoded message as a new slice
func uuencode(dec_msg []byte) (enc_msg []byte) {
	length := byte(len(dec_msg))
	dec_len := length
	mod := dec_len % 3

	if mod != 0 {
		dec_len += 3 - mod
		for i := length; i < dec_len; i++ {
			dec_msg[i] = 0
		}
	}

	for i := byte(0); i < dec_len-2; i += 3 {
		enc_msg = append(enc_msg, (dec_msg[i+0]&0x3F)+32)
		enc_msg = append(enc_msg, (dec_msg[i+1]&0x3F)+32)
		enc_msg = append(enc_msg, (dec_msg[i+2]&0x3F)+32)

		enc_byte := byte(0)

		if (dec_msg[i+0] & 0x40) != 0 {
			enc_byte |= 0x01
		}
		if (dec_msg[i+0] & 0x80) != 0 {
			enc_byte |= 0x02
		}
		if (dec_msg[i+1] & 0x40) != 0 {
			enc_byte |= 0x04
		}
		if (dec_msg[i+1] & 0x80) != 0 {
			enc_byte |= 0x08
		}
		if (dec_msg[i+2] & 0x40) != 0 {
			enc_byte |= 0x10
		}
		if (dec_msg[i+2] & 0x80) != 0 {
			enc_byte |= 0x20
		}

		enc_msg = append(enc_msg, enc_byte+32)
	}

	return
}

// Decodes a message and returns the decoded message as a new slice
func uudecode(enc_msg []byte) (dec_msg []byte) {
	length := byte(len(enc_msg))
	enc_len := length

	if (enc_len % 4) != 0 {
		length = 0
		return
	}

	for i := byte(0); i < enc_len/4; i++ {
		dec_msg = append(dec_msg, enc_msg[(i*4)+0]-32)
		dec_msg = append(dec_msg, enc_msg[(i*4)+1]-32)
		dec_msg = append(dec_msg, enc_msg[(i*4)+2]-32)

		enc_byte := enc_msg[(i*4)+3] - 32

		if (enc_byte & 0x01) != 0 {
			dec_msg[(i*3)+0] |= 0x40
		}
		if (enc_byte & 0x02) != 0 {
			dec_msg[(i*3)+0] |= 0x80
		}
		if (enc_byte & 0x04) != 0 {
			dec_msg[(i*3)+1] |= 0x40
		}
		if (enc_byte & 0x08) != 0 {
			dec_msg[(i*3)+1] |= 0x80
		}
		if (enc_byte & 0x10) != 0 {
			dec_msg[(i*3)+2] |= 0x40
		}
		if (enc_byte & 0x20) != 0 {
			dec_msg[(i*3)+2] |= 0x80
		}
	}

	return
}

// Function to handle if this device should respond to a message
// pkt    : Packet to check
// return : true if it is for this device
func (p *PCI) intendedRecepient(pkt PCIPacket) bool {
	// Function needs to exist to handle "for-all" type messages (0xF)
	return ((pkt.destination_ID == p.id) || (pkt.destination_ID == 0xF))
}

func (p *PCI) hasFreeChannel() bool {
	return p.getFreeChannel() < 0
}

func (p *PCI) getFreeChannel() int {
	for i := 0; i < len(p.channels); i++ {
		if p.channels[i].group == GROUP_NONE {
			return i
		}
	}

	return -1
}

// Function to return index of a channel of the PCI object that responds to a group
// Returns -1 if no channel is assigned to that group
func (p *PCI) getChannelFromGroup(group byte) int {
	for i, item := range p.channels {
		if item.group == group {
			return i
		}
	}

	return -1
}
