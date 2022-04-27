package main

import (
	"errors"
	"fmt"
)

func main() {
	encPacket := []byte{0xA1, 0x49, 0x20, 0x22, 0x20, 0x38, 0x20, 0x24, 0x0D}
	length := byte(len(encPacket) - 1)

	fmt.Println("Encoded:", encPacket)

	decStream := uudecode(encPacket, &length)

	pkt, err := messageToPacket(decStream)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Decoded:", pkt)
}

type PCIPacket struct {
	source_ID      uint8
	destination_ID uint8
	command        PCI_CMD
	parameters     []byte
}

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

// messageToPacket returns a PCIPacket object from a decoded byte slice
func messageToPacket(decodedMsg []byte) (pkt PCIPacket, err error) {
	// Handle byte 0 (S_ID | D_ID)
	pkt.source_ID = decodedMsg[0] >> 4
	pkt.destination_ID = decodedMsg[0] & 0x0F
	// Handle byte 1 (LEN | CMD)
	length := decodedMsg[1] >> 3
	pkt.command = PCI_CMD(decodedMsg[1] & 0x07)
	// Get parameters
	for i := byte(0); i < length-3; i++ {
		pkt.parameters = append(pkt.parameters, decodedMsg[i+2])
	}
	// Get checksum
	if calc_chk(decodedMsg, length, 1) != 0 {
		err = errors.New("checksum failed")
	}

	return
}

// packetToMessage returns a byte slice ready to encode and transmit via serial
// from the calling object's specifications
func (p *PCIPacket) packetToMessage() (decodedMsg []byte) {
	// Handle byte 0 (S_ID | D_ID)
	decodedMsg = append(decodedMsg, ((p.source_ID << 4) | p.destination_ID))
	// Handle byte 1 (Len | Cmd)
	length := p.calcLength()
	decodedMsg = append(decodedMsg, ((length << 3) | byte(p.command)))
	// Handle parameters
	for i := 0; i < len(p.parameters); i++ {
		decodedMsg = append(decodedMsg, p.parameters[i])
	}
	// Handle checksum
	decodedMsg = append(decodedMsg, calc_chk(decodedMsg, length-1, 0))

	// Return
	return
}

// calcLength returns the message length in bytes of the calling object's packet
func (p *PCIPacket) calcLength() (messageLength byte) {
	messageLength = 0
	// Add bytes for header (S_ID | D_ID) (LEN | CMD)
	messageLength += 2
	// Add number of parameters
	messageLength += byte(len(p.parameters))
	// Add byte for checksum
	messageLength += 1
	return
}

// calc_chk returns the calculated checksum of a decoded message
// Mode requires the input of either read(1) or write (0)
// In read mode, the final checksum should be 0, referring to a good packet
// In write mode, the final checksum will be the checksum that should be transmitted
func calc_chk(msg []byte, len byte, mode byte) byte {
	chk := int(0)

	if mode == 1 {
		len--
	}

	for i := int(0); i < int(len); i++ {
		chk += int(msg[i])

		if (chk & 0x0100) != 0 {
			chk = (chk & 0xFF) | 0x01
		} else {
			chk &= ^0x01
		}
	}

	if mode == 0 {
		// Calculate checksum for transmit
		return (1 + ^(byte(chk)))
	} else {
		// Calculate output checksum to ensure successful reception
		return byte(chk + int(msg[len]))
	}
}

func uudecode(enc_msg []byte, len *byte) (dec_msg []byte) {
	enc_len := *len

	if (enc_len % 4) != 0 {
		*len = 0
		return
	} else {
		*len = (enc_len / 4) * 3
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