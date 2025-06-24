package set

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/types"
	"maps"
)

type Set[E comparable] map[E]types.Unit

func New[E comparable](cap int) Set[E] {
	return make(Set[E], cap)
}

func Of[E comparable](es ...E) Set[E] {
	s := make(map[E]types.Unit, len(es))
	for _, e := range es {
		s[e] = types.Unit{}
	}
	return s
}

func OfEmpty[E comparable]() Set[E] {
	return make(Set[E], 0)
}

func (s Set[E]) ForEach(fn func(E)) {
	for e := range s {
		fn(e)
	}
}

func (s Set[E]) Len() int {
	return len(s)
}

func (s Set[E]) Empty() bool {
	return len(s) == 0
}

func (s Set[E]) Insert(element E) {
	s[element] = types.Unit{}
}

func (s Set[E]) Inserts(elements vec.Vec[E]) {
	for _, element := range elements {
		s[element] = types.Unit{}
	}
}
func (s Set[E]) InsertSelf(element E) Set[E] {
	s[element] = types.Unit{}
	return s
}

func (s Set[E]) Remove(element E) {
	delete(s, element)
}

func (s Set[E]) Contains(element E) bool {
	_, ok := s[element]
	return ok
}

func (s Set[E]) Or(other Set[E]) Set[E] {
	s0, s1 := s, other
	// Select a larger Set[E] as a baseline
	if s0.Len() < s1.Len() {
		s0, s1 = s1, s0
	}
	s_ := maps.Clone(s0)
	for e := range s1 {
		s_[e] = types.Unit{}
	}
	return s_
}

func (s Set[E]) And(other Set[E]) Set[E] {
	s0, s1 := s, other
	// Choose a smaller Set[E] as the baseline for iteration
	if s0.Len() < s1.Len() {
		s0, s1 = s1, s0
	}
	s_ := New[E]((s1.Len() + 1) / 2)
	for e := range s1 {
		if s0.Contains(e) {
			s_[e] = types.Unit{}
		}
	}
	return s_
}

func (s Set[E]) Sub(other Set[E]) Set[E] {
	if s.Len() < other.Len()*2 {
		// When the Set[E] is small, check them one by one
		s_ := New[E]((s.Len() + 1) / 2)
		for e := range s {
			if !other.Contains(e) {
				s_[e] = types.Unit{}
			}
		}
		return s_
	} else {
		// Clone the current Set[E] and remove elements from another Set[E]
		s_ := maps.Clone(s)
		for e := range other {
			s_.Remove(e)
		}
		return s_
	}
}

func (s Set[E]) Xor(other Set[E]) Set[E] {
	s0, s1 := s, other
	// Select a larger Set[E] as a baseline
	if s0.Len() < s1.Len() {
		s0, s1 = s1, s0
	}
	s_ := maps.Clone(s0)
	for e := range s1 {
		if s_.Contains(e) {
			s_.Remove(e)
		} else {
			s_[e] = types.Unit{}
		}
	}
	return s_
}

func (s Set[E]) ToVec() vec.Vec[E] {
	v := vec.New[E](len(s))
	for e := range s {
		v.Append(e)
	}
	return v
}

func (s Set[E]) Clear() {
	clear(s)
}

func (s Set[E]) String() string {
	return fmt.Sprintf("set%v", s.ToVec())
}

// MarshalJSON Implement JSON serialization
func (s Set[E]) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(s.ToVec())
}

// UnmarshalJSON Implement JSON deserialization
func (s *Set[E]) UnmarshalJSON(data []byte) error {
	arr := new(vec.Vec[E])
	err := sonic.Unmarshal(data, arr)
	if err == nil {
		*s = Of[E](*arr...)
	}
	return err
}
