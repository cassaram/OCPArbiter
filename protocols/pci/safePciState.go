package pci

import (
	"fmt"
	"sync"
)

type safePciState struct {
	state PCI_CONNECTION_STATE
	mute  sync.RWMutex
}

func (s *safePciState) Get() PCI_CONNECTION_STATE {
	s.mute.Lock()
	defer s.mute.Unlock()

	return s.state
}

func (s *safePciState) Set(state PCI_CONNECTION_STATE) {
	s.mute.Lock()
	defer s.mute.Unlock()

	s.state = state
	fmt.Println("New state:", s.state)
}
