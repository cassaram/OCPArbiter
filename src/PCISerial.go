package main

import (
	"log"

	"github.com/tarm/serial"
)

// Implements a PC / PCI like serial interface at level 0

type PCISerial struct {
	receiveBuffer []byte
	serialPort    *serial.Port
	isRunning     bool
}

// UART initialisation, receive buffer initialisation
func (s *PCISerial) ser_init(portName string) {
	serialConfig := &serial.Config{Name: portName, Baud: 9600}
	port, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Fatal(err)
	}
	s.serialPort = port

	// Setup receive buffer
	s.receiveBuffer = make([]byte, 128)

	// Inform dependent functions
	s.isRunning = true
}

// Send a character on the serial link
func (s *PCISerial) putser(c rune) {

}

// Check if a character is available within the receive buffer
func (s *PCISerial) chkser() bool {
	return false
}

// Update receiver buffer
func (s *PCISerial) read() {
	if s.isRunning {
		_, err := s.serialPort.Read(s.receiveBuffer)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Take a character from the receive buffer
func (s *PCISerial) getser() byte {
	// Update receiever buffer
	s.read()

}
