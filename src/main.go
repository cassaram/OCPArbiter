package main

import (
	"github.com/cassaram/ocparbiter/src/gvocp"
)

func main() {
	var serialTest gvocp.PCISerial
	serialTest.ser_init("COM3")
}
