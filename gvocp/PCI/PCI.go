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

const (
	GROUP_NONE byte = 0x10
)

type PCI struct {
	port                serial.Port
	rxBuffer            []byte
	dataMessageHandler  func(PCIPacket)
	testRingInitialized bool
	d_id                byte
	group               byte
	grpChannels         []byte
}

func (p *PCI) setPort(name string, destinationID byte) {
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
	p.d_id = destinationID
	p.group = GROUP_NONE
}

// Function to send a PCIPacket
func (p *PCI) sendPacket(pkt PCIPacket) (err error) {
	// Encode into byte slice
	dec_msg := pkt.packetToMessage()
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
func (p *PCI) handleData() {
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
			// Decode message
			dec_msg := uudecode(enc_msg)
			// Convert to PCI Packet
			pkt, err := messageToPacket(dec_msg)
			if err != nil {
				log.Printf("Malformed packet received. Encoded: %[1]v, Decoded: %[2]v", enc_msg, dec_msg)
			}
			// Announce that packet arrived
			p.packetArrived(pkt)
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
	switch pkt.command {
	case DATA_MESSAGE:
		p.dataMessageHandler(pkt)
	case TEST_RING:
		p.handleTestRing(pkt)
	case ASSIGN_TO_GROUP:

	}
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

	// Ensure ring has already been reset
	if !p.testRingInitialized {
		// Set reset request bit
		flag |= 0x01
	}

	// Handle reset request
	if resetRing {
		// Reset all channel assignments
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
}

// Handles ASSIGN_TO_GROUP messages
func (p *PCI) handleAssignToGroup(pkt PCIPacket) {
	grpByte := pkt.parameters[0]
	grpNum := grpByte & 0x0F
	//noInit := ((grpByte & 0x80) != 0)

	if (!p.intendedRecepient(pkt)) || (p.group == GROUP_NONE) {
		// Cannot execute command, pass along instead
		p.sendPacket(pkt)
		return
	}

	// Connect & send CONNECTED_TO_PART command
	p.group = grpNum

	stat := byte(0x02) // Channel # (7-4), Camera (1), Base station (0)

	var connectedParams []byte
	connectedParams = append(connectedParams, grpNum, stat)

	connectedPkt := PCIPacket{
		source_ID:      p.d_id,
		destination_ID: pkt.source_ID,
		command:        CONNECTED_TO_PART,
		parameters:     connectedParams,
	}

	p.sendPacket(connectedPkt)
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
	grp := pkt.parameters[0]
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

	for i := byte(0); i < dec_len; i += 3 {
		enc_msg = append(enc_msg, (dec_msg[(i*3)+0]&0x3F)+32)
		enc_msg = append(enc_msg, (dec_msg[(i*3)+1]&0x3F)+32)
		enc_msg = append(enc_msg, (dec_msg[(i*3)+2]&0x3F)+32)

		enc_byte := byte(0)

		if (dec_msg[(i*3)+0] & 0x40) != 0 {
			enc_byte |= 0x01
		}
		if (dec_msg[(i*3)+0] & 0x80) != 0 {
			enc_byte |= 0x02
		}
		if (dec_msg[(i*3)+1] & 0x40) != 0 {
			enc_byte |= 0x04
		}
		if (dec_msg[(i*3)+1] & 0x80) != 0 {
			enc_byte |= 0x08
		}
		if (dec_msg[(i*3)+2] & 0x40) != 0 {
			enc_byte |= 0x10
		}
		if (dec_msg[(i*3)+2] & 0x80) != 0 {
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
	return ((pkt.destination_ID == p.d_id) || (pkt.destination_ID == 0xF))
}
