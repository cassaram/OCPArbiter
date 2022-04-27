package pci

import (
	"log"

	"go.bug.st/serial"
)

const (
	PCI_REC_IDLE  int = 1
	PCI_REC_READY int = 2
	PCI_REC_ERROR int = 3
)

type PCI struct {
	port     serial.Port
	rxBuffer []byte
	rec_buf  []byte
	rec_cnt  uint32
	rec_err  uint32
	rec_rdy  uint32
	rec_max  uint32
}

func (p *PCI) setPort(name string) {
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
	p.rxBuffer = make([]byte, 100)
}

/***************************************************************************
;
; Name    : uuencode
; Input   : byte* dec_msg: source
;           byte* enc_msg: destination
;           byte* len    : length
; Output  : <len> will contain the new encoded packet length
; Function: Encodes a PCI packet <dec_msg> with length <len> to a
;           (pseudo) uuencode encoded packet in <enc_msg>
; Ref.    :
;
;***************************************************************************/
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

/***************************************************************************
;
; Name    : uudecode
; Input   : byte* enc_msg: source
;           byte* dec_msg: destination
;           byte* len    : length
; Output  : <len> will contain the new decoded packet length
; Function: Decodes a (pseudo) uuencode encoded PCI packet <enc_msg>
;           with length <len> to a decoded packet in <dec_msg>
; Ref.    :
;
;***************************************************************************/
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

/***************************************************************************
;
; Name    : chk_msg_state
; Input   : byte* cmd
; Output  : int: PCI_REC_IDLE is returned when no (complete) message is available
;                PCI_REC_READY is returned when a complete message is available
;                PCI_REC_ERROR on format error
; Function: Checks the state of eventual received PCI messages
; Ref.    :
;
;***************************************************************************/
func (p *PCI) chk_msg_state(cmd *byte) int {
	var c byte

	p.rec_rdy = 0

	for len(p.rxBuffer) != 0 {
		c = p.rxBuffer[0]

		// Check if packet is a header
		if c&0x80 != 0 {
			// Packet Header
			if p.rec_cnt != 0 {
				// Format error in header
				p.rec_cnt = 0
				p.rec_err = 1
				return PCI_REC_ERROR
			} else {
				// Valid packet header
				// Strip header bit
				c &= 0x7F
				p.rec_err = 0
			}
		} else {
			// Check if header expected
			if p.rec_cnt == 0 {
				// Check error state
				if p.rec_err == 0 {
					p.rec_err = 1
					return PCI_REC_ERROR
				} else {
					return PCI_REC_IDLE
				}
			}
		}

		// Check for packet terminator
		if c == 13 {
			*cmd = (p.rxBuffer[1] - 32&0x07)

			p.rec_rdy = 1
			return PCI_REC_READY
		}

		// Check if we have space in buffer
		if p.rec_cnt < p.rec_max {
			// Sore in buffer
			p.rxBuffer[p.rec_cnt] = c
			p.rec_cnt++
		} else {
			// Packet too long
			p.rec_cnt = 0
			p.rec_err = 1
			return PCI_REC_ERROR
		}
	}

	return PCI_REC_IDLE
}
