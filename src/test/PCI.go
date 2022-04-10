package main

import "fmt"

// Implements the layer 1 protocol

func uudecode(enc_msg []byte, dec_msg []byte, len *byte) {
	enc_len := *len

	if (enc_len % 4) != 0 {
		*len = 0
		return
	} else {
		*len = (enc_len / 4) * 3
	}

	for i := byte(0); i < enc_len/4; i++ {
		dec_msg[(i*3)+0] = enc_msg[(i*4)+0] - 32
		dec_msg[(i*3)+1] = enc_msg[(i*4)+1] - 32
		dec_msg[(i*3)+2] = enc_msg[(i*4)+2] - 32

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
}

func main() {
	testData := []byte{
		0xFF,
		0xF7,
		0xFD,
		0xFE,
		0x00,
		0x7F,
		0x82,
		0x00,
		0xFE,
		0xFE,
		0xFF,
	}
	testLength := byte(8)
	testDecode := make([]byte, len(testData))

	fmt.Println("Encoded:", testData)

	uudecode(testData, testDecode, &testLength)

	fmt.Println("Decoded:", testDecode)
}
