package main

import (
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

type OCP400 struct {
	PortName string
	Port     io.ReadWriteCloser
}

func (o *OCP400) connectSerial() {
	// Configure options for serial port
	serialOptions := serial.OpenOptions{
		PortName: o.PortName,
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
	}

	// Open port
	port, err := serial.Open(serialOptions)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	o.Port = port
}
