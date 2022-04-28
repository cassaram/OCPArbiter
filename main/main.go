package main

import (
	"fmt"

	pci "github.com/cassaram/ocparbiter/gvocp/PCI"
)

func listenLoop(p *pci.PCI) {
	fmt.Println("Listen Loop Called")
	for {
		p.HandleData()
	}
}

func main() {
	var p pci.PCI
	p.SetPort("COM1", 1)

	listenLoop(&p)

}
