package pci

import "sync"

type pciChannel struct {
	group        uint8
	controllerID uint8
}

type safePciChannels struct {
	channels [4]pciChannel
	mute     sync.Mutex
}

func (s *safePciChannels) NewGroup(group uint8, controllerID uint8) int8 {
	s.mute.Lock()
	defer s.mute.Unlock()

	for i := 0; i < len(s.channels); i++ {
		if s.channels[i].group == 0 && s.channels[i].controllerID == 0 {
			// assign channel
			s.channels[i].group = group
			s.channels[i].controllerID = controllerID

			// Return index
			return int8(i)
		}
	}

	// return failure
	return -1
}

func (s *safePciChannels) GetController(group uint8) int8 {
	s.mute.Lock()
	defer s.mute.Unlock()

	for _, val := range s.channels {
		if val.group == group {
			return int8(val.controllerID)
		}
	}

	return -1
}

func (s *safePciChannels) GetChannelID(group uint8) int8 {
	s.mute.Lock()
	defer s.mute.Unlock()

	for i, val := range s.channels {
		if val.group == group {
			return int8(i)
		}
	}

	return -1
}

func (s *safePciChannels) DeleteGroup(group uint8) {
	s.mute.Lock()
	defer s.mute.Unlock()

	for _, val := range s.channels {
		if val.group == group {
			val.group = 0
			val.controllerID = 0
		}
	}
}

func (s *safePciChannels) ResetChannels() {
	s.mute.Lock()
	defer s.mute.Unlock()

	for _, val := range s.channels {
		val.group = 0
		val.controllerID = 0
	}
}
