package set

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/types"
	"sync"
)

type SyncSet[E comparable] struct {
	rw sync.RWMutex
	v  map[E]types.Unit
}

func SyncNew[E comparable](cap int) SyncSet[E] {
	return SyncSet[E]{v: make(map[E]types.Unit, cap)}
}

func SyncOf[E comparable](es ...E) SyncSet[E] {
	s := make(map[E]types.Unit, len(es))
	for _, e := range es {
		s[e] = types.Unit{}
	}
	return SyncSet[E]{v: s}
}

func SyncOfEmpty[E comparable]() SyncSet[E] {
	return SyncNew[E](0)
}

func (s *SyncSet[E]) ForEach(fn func(E)) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	for e := range s.v {
		fn(e)
	}
}

func (s *SyncSet[E]) Len() int {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return len(s.v)
}

func (s *SyncSet[E]) Empty() bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return len(s.v) == 0
}

func (s *SyncSet[E]) Insert(element E) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.v[element] = types.Unit{}
}

func (s *SyncSet[E]) Inserts(elements vec.Vec[E]) {
	s.rw.Lock()
	defer s.rw.Unlock()
	for _, element := range elements {
		s.v[element] = types.Unit{}
	}
}
func (s *SyncSet[E]) InsertSelf(element E) *SyncSet[E] {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.v[element] = types.Unit{}
	return s
}

func (s *SyncSet[E]) Remove(element E) {
	s.rw.Lock()
	defer s.rw.Unlock()
	delete(s.v, element)
}

func (s *SyncSet[E]) Contains(element E) bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	_, ok := s.v[element]
	return ok
}

func (s *SyncSet[E]) contains(element E) bool {
	_, ok := s.v[element]
	return ok
}

func (s *SyncSet[E]) Or(other *SyncSet[E]) *SyncSet[E] {
	s.rw.RLock()
	other.rw.RLock()
	defer s.rw.RUnlock()
	defer other.rw.RUnlock()
	s0, s1 := s, other
	// Select a larger Set[E] as a baseline
	if len(s0.v) < len(s1.v) {
		s0, s1 = s1, s0
	}
	s_ := &SyncSet[E]{v: s0.v}
	for e := range s1.v {
		s_.v[e] = types.Unit{}
	}
	return s_
}

func (s *SyncSet[E]) And(other *SyncSet[E]) *SyncSet[E] {
	s.rw.RLock()
	other.rw.RLock()
	defer s.rw.RUnlock()
	defer other.rw.RUnlock()
	s0, s1 := s, other
	// Choose a smaller Set[E] as the baseline for iteration
	length := len(s1.v)
	if len(s0.v) < length {
		s0, s1 = s1, s0
	}
	s_ := SyncNew[E]((length + 1) / 2)
	for e := range s1.v {
		if s0.contains(e) {
			s_.v[e] = types.Unit{}
		}
	}
	return &s_
}

func (s *SyncSet[E]) Sub(other *SyncSet[E]) *SyncSet[E] {
	s.rw.RLock()
	other.rw.RLock()
	defer s.rw.RUnlock()
	defer other.rw.RUnlock()
	length := len(s.v)
	if length < len(other.v)*2 {
		// When the Set[E] is small, check them one by one
		s_ := SyncNew[E]((length + 1) / 2)
		for e := range s.v {
			if !other.contains(e) {
				s_.v[e] = types.Unit{}
			}
		}
		return &s_
	} else {
		// Clone the current Set[E] and remove elements from another Set[E]
		s_ := &SyncSet[E]{v: s.v}
		for e := range other.v {
			s_.Remove(e)
		}
		return s_
	}
}

func (s *SyncSet[E]) Xor(other *SyncSet[E]) *SyncSet[E] {
	s.rw.RLock()
	other.rw.RLock()
	defer s.rw.RUnlock()
	defer other.rw.RUnlock()
	s0, s1 := s, other
	// Select a larger Set[E] as a baseline
	if len(s0.v) < len(s1.v) {
		s0, s1 = s1, s0
	}
	s_ := &SyncSet[E]{v: s0.v}
	for e := range s1.v {
		if s_.contains(e) {
			s_.Remove(e)
		} else {
			s_.v[e] = types.Unit{}
		}
	}
	return s_
}

func (s *SyncSet[E]) ToVec() vec.Vec[E] {
	s.rw.RLock()
	defer s.rw.RUnlock()
	v := vec.New[E](len(s.v))
	for e := range s.v {
		v.Append(e)
	}
	return v
}

func (s *SyncSet[E]) toVec() vec.Vec[E] {
	v := vec.New[E](len(s.v))
	for e := range s.v {
		v.Append(e)
	}
	return v
}

func (s *SyncSet[E]) Clear() {
	s.rw.Lock()
	defer s.rw.Unlock()
	clear(s.v)
}

func (s *SyncSet[E]) String() string {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return fmt.Sprintf("set%v", s.toVec())
}

// MarshalJSON Implement JSON serialization
func (s *SyncSet[E]) MarshalJSON() ([]byte, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return sonic.Marshal(s.toVec())
}

// UnmarshalJSON Implement JSON deserialization
func (s *SyncSet[E]) UnmarshalJSON(data []byte) error {
	s.rw.Lock()
	defer s.rw.Unlock()
	arr := new(vec.Vec[E])
	err := sonic.Unmarshal(data, arr)
	if err == nil {
		*s = SyncOf[E](*arr...)
	}
	return err
}
