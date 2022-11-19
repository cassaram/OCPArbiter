package settings

import "sync"

type SafeSettings struct {
	data map[string]Setting
	mute sync.Mutex
}

func (s *SafeSettings) GetLength() int {
	s.mute.Lock()
	defer s.mute.Unlock()

	return len(s.data)
}

func (s *SafeSettings) IsEmpty() bool {
	return s.GetLength() == 0
}

func (s *SafeSettings) Add(set Setting) {
	s.mute.Lock()
	defer s.mute.Unlock()

	s.data[set.Id] = set
}

func (s *SafeSettings) Remove(id string) {
	s.mute.Lock()
	defer s.mute.Unlock()

	delete(s.data, id)
}

func (s *SafeSettings) Get(id string) string {
	s.mute.Lock()
	defer s.mute.Unlock()

	return s.data[id].Value
}

func (s *SafeSettings) GetInt(id string) int {
	s.mute.Lock()
	defer s.mute.Unlock()

	set := s.data[id]

	return set.ValueInt()
}

func (s *SafeSettings) Change(id string, val string) {
	s.mute.Lock()
	defer s.mute.Unlock()

	set := s.data[id]
	set.Value = val
}

func (s *SafeSettings) ResetDefault(id string) {
	s.mute.Lock()
	defer s.mute.Unlock()

	set := s.data[id]
	set.Value = set.Default
}

func (s *SafeSettings) Export() []Setting {
	s.mute.Lock()
	defer s.mute.Unlock()

	result := make([]Setting, s.GetLength())

	i := 0
	for _, val := range s.data {
		result[i] = val
		i++
	}

	return result
}
