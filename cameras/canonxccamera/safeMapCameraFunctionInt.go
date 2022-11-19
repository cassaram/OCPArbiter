package canonxccamera

import (
	"sync"

	"github.com/cassaram/ocparbiter/common"
)

type safeMapCameraFunctionInt struct {
	data map[common.CameraFunction]int
	mute sync.Mutex
}

func newSafeMapCameraFunctionInt() *safeMapCameraFunctionInt {
	return &safeMapCameraFunctionInt{
		data: map[common.CameraFunction]int{},
	}
}

func (s *safeMapCameraFunctionInt) Set(key common.CameraFunction, value int) {
	s.mute.Lock()
	defer s.mute.Unlock()

	s.data[key] = value
}

func (s *safeMapCameraFunctionInt) Get(key common.CameraFunction) (int, bool) {
	s.mute.Lock()
	defer s.mute.Unlock()

	val, ok := s.data[key]
	return val, ok
}

func (s *safeMapCameraFunctionInt) Delete(key common.CameraFunction) {
	s.mute.Lock()
	defer s.mute.Unlock()

	delete(s.data, key)
}
