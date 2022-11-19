package common

import "sync"

type DeviceState uint8

const (
	CTRL_AWAIT_START DeviceState = iota
	CTRL_CONNECTING
	CTRL_CONNECTED
)

type SafeDeviceState struct {
	state DeviceState
	mute  sync.Mutex
}

func (s *SafeDeviceState) Get() DeviceState {
	s.mute.Lock()
	defer s.mute.Unlock()

	return s.state
}

func (s *SafeDeviceState) Set(state DeviceState) {
	s.mute.Lock()
	defer s.mute.Unlock()

	s.state = state
}
