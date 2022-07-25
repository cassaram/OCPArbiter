package testcam

import (
	"sync"

	"github.com/cassaram/ocparbiter/common"
)

type safeCameraFunctionInt struct {
	data map[common.CameraFunction]int
	mute sync.Mutex
}

func newSafeCameraFunctionInt() *safeCameraFunctionInt {
	return &safeCameraFunctionInt{
		data: map[common.CameraFunction]int{},
	}
}

func (s *safeCameraFunctionInt) Set(key common.CameraFunction, value int) {
	s.mute.Lock()
	defer s.mute.Unlock()

	s.data[key] = value
}

func (s *safeCameraFunctionInt) Get(key common.CameraFunction) (int, bool) {
	s.mute.Lock()
	defer s.mute.Unlock()

	val, ok := s.data[key]
	return val, ok
}

func (s *safeCameraFunctionInt) Delete(key common.CameraFunction) {
	s.mute.Lock()
	defer s.mute.Unlock()

	delete(s.data, key)
}
