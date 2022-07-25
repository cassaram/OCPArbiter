package pci

import (
	"sync"
	"time"
)

type safeTime struct {
	time time.Time
	mute sync.Mutex
}

func (s *safeTime) GetTime() time.Time {
	s.mute.Lock()
	defer s.mute.Unlock()

	return s.time
}

func (s *safeTime) GetTimeSince() time.Duration {
	s.mute.Lock()
	defer s.mute.Unlock()

	return time.Since(s.time)
}

func (s *safeTime) SetNow() {
	s.mute.Lock()
	defer s.mute.Unlock()

	s.time = time.Now()
}
