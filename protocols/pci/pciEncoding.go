package pci

// Performs uuencoding in compliance with the PCI protocol
func uuencode(src []byte) []byte {
	// Copy to new message
	dec_msg := make([]byte, len(src))
	copy(dec_msg, src)

	// Append zeroes until message length is a multiple of 3
	length := len(dec_msg)
	dec_len := length
	mod := dec_len % 3

	if mod != 0 {
		dec_len += 3 - mod
		for i := length; i < dec_len; i++ {
			dec_msg[i] = 0
		}
	}

	// Create encoded message slice
	enc_len := ((dec_len / 3) * 4)
	enc_msg := make([]byte, enc_len)

	// handle encode
	for i := 0; i < dec_len/3; i++ {
		// Map bits 5-0 onto the encoded message w/ ASCII (+32) conversion
		enc_msg[(i*4)+0] = (dec_msg[(i*3)+0] & 0x3F) + 32
		enc_msg[(i*4)+1] = (dec_msg[(i*3)+1] & 0x3F) + 32
		enc_msg[(i*4)+2] = (dec_msg[(i*3)+2] & 0x3F) + 32

		// Map bits 7-6 onto the next byte
		if (dec_msg[(i*3)+0] & 0x40) != 0 {
			enc_msg[(i*4)+3] |= 0x01
		}
		if (dec_msg[(i*3)+0] & 0x80) != 0 {
			enc_msg[(i*4)+3] |= 0x02
		}
		if (dec_msg[(i*3)+1] & 0x40) != 0 {
			enc_msg[(i*4)+3] |= 0x04
		}
		if (dec_msg[(i*3)+1] & 0x80) != 0 {
			enc_msg[(i*4)+3] |= 0x08
		}
		if (dec_msg[(i*3)+2] & 0x40) != 0 {
			enc_msg[(i*4)+3] |= 0x10
		}
		if (dec_msg[(i*3)+2] & 0x80) != 0 {
			enc_msg[(i*4)+3] |= 0x20
		}

		// Perform ASCII (+32) conversion on new byte
		enc_msg[(i*4)+3] += 32
	}

	// Return uuencoded message
	return enc_msg
}

// Performs uudecoding in compliance with the PCI Protocol
func uudecode(src []byte) []byte {
	// Copy to new message
	enc_msg := make([]byte, len(src))
	copy(enc_msg, src)

	// Ensure length is correct before decode attempt
	length := len(enc_msg)
	enc_len := length
	if (enc_len % 4) != 0 {
		// Packet does not conform
		// Return empty slice
		return make([]byte, 0)
	}

	// Make decoded byte slice
	dec_len := (enc_len / 4) * 3
	dec_msg := make([]byte, dec_len)

	for i := 0; i < dec_len/3; i++ {
		// Remove ASCII conversion and extract bits
		dec_msg[(i*3)+0] = (enc_msg[(i*4)+0]) - 32
		dec_msg[(i*3)+1] = (enc_msg[(i*4)+1]) - 32
		dec_msg[(i*3)+2] = (enc_msg[(i*4)+2]) - 32

		// Remove ASII conevsion and extract bits from extra byte
		extraByte := enc_msg[(i*4)+3] - 32
		if (extraByte & 0x01) != 0 {
			dec_msg[(i*3)+0] |= 0x40
		}
		if (extraByte & 0x02) != 0 {
			dec_msg[(i*3)+0] |= 0x80
		}
		if (extraByte & 0x04) != 0 {
			dec_msg[(i*3)+1] |= 0x40
		}
		if (extraByte & 0x08) != 0 {
			dec_msg[(i*3)+1] |= 0x80
		}
		if (extraByte & 0x10) != 0 {
			dec_msg[(i*3)+2] |= 0x40
		}
		if (extraByte & 0x20) != 0 {
			dec_msg[(i*3)+2] |= 0x80
		}
	}

	// Return decoded message
	return dec_msg
}
