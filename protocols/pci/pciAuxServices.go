package pci

import (
	"sync"
	"time"
)

type safeTestRingTime struct {
	data time.Time
	mute sync.Mutex
}

func (s *safeTestRingTime) GetDelta() float64 {
	s.mute.Lock()
	defer s.mute.Unlock()

	delta := time.Now().Sub(s.data)

	return delta.Seconds()
}

func (s *safeTestRingTime) Update() {
	s.mute.Lock()
	defer s.mute.Unlock()

	s.data = time.Now()
}
