package main

import "fmt"

func main() {
	raw_data := []byte{0xA1, 0x49, 0x20, 0x22, 0x20, 0x38, 0x20, 0x24, 0x0D}
	// Remove tail
	raw_data = raw_data[:8]
	fmt.Printf("Encoded data (raw):           [% x]\n", raw_data)
	dec_raw := uudecode(raw_data)
	raw_data[0] &= 0x7F
	fmt.Printf("Encoded data (no header bit): [% x]\n", raw_data)
	dec_header := uudecode(raw_data)
	fmt.Printf("Decoded data (raw):           [% x]\n", dec_raw)
	fmt.Printf("Decoded data (no header bit): [% x]\n", dec_header)

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

	for i := byte(0); i < dec_len/3; i++ {
		enc_msg = append(enc_msg, (dec_msg[(i*4)+0]&0x3F)+32)
		enc_msg = append(enc_msg, (dec_msg[(i*4)+1]&0x3F)+32)
		enc_msg = append(enc_msg, (dec_msg[(i*4)+2]&0x3F)+32)

		enc_byte := byte(0)

		if (dec_msg[(i*4)+0] & 0x40) != 0 {
			enc_byte |= 0x01
		}
		if (dec_msg[(i*4)+0] & 0x80) != 0 {
			enc_byte |= 0x02
		}
		if (dec_msg[(i*4)+1] & 0x40) != 0 {
			enc_byte |= 0x04
		}
		if (dec_msg[(i*4)+1] & 0x80) != 0 {
			enc_byte |= 0x08
		}
		if (dec_msg[(i*4)+2] & 0x40) != 0 {
			enc_byte |= 0x10
		}
		if (dec_msg[(i*4)+2] & 0x80) != 0 {
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
