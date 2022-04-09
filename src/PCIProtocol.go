package main

type PCIPacket struct {
	S_ID   uint8
	D_ID   uint8
	CNT    uint8
	CMD    uint8
	LEN    uint16
	Params []uint8
	CHK    uint8
}

func (m *PCIPacket) parsePacket(s []uint8) {
	index := 0
	m.S_ID = (s[index] >> 4)
	m.D_ID = (s[index] & 15)
	index++
	m.CNT = (s[index] >> 3)
	m.CMD = (s[index] & 7)
	index++
	// Check count for how many length bytes we have
	if m.CNT >= 3 {
		// Count is length of message
		m.LEN = uint16(m.CNT)
	} else if m.CNT == 0 {
		// LEN0 is added to contain length
		m.LEN = uint16(s[index])
	} else if m.CNT == 1 {
		// LEN0 and LEN1 are added to contain length
		m.LEN = uint16(s[index])
		index++
		m.LEN = m.LEN | uint16(s[index])
		index++
	}
	for i := uint16(0); i < m.LEN; i++ {

	}

}
