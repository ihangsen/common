package vec

import (
	"github.com/ihangsen/common/src/utils/option"
	"slices"
	"sync"
)

type SyncVec[E any] struct {
	rw sync.RWMutex
	v  []E
}

func SyncNew[E any](cap int) SyncVec[E] {
	return SyncVec[E]{v: make([]E, 0, cap)}
}

func SyncOf[E any](arr ...E) SyncVec[E] {
	return SyncVec[E]{v: arr}
}

func SyncOfEmpty[E any]() SyncVec[E] {
	return SyncNew[E](0)
}

func (v *SyncVec[E]) ToVec() Vec[E] {
	return v.v
}

func (v *SyncVec[E]) ForEach(fn func(E)) {
	v.rw.RLock()
	defer v.rw.RUnlock()
	for _, e := range v.v {
		fn(e)
	}
}

func (v *SyncVec[E]) Len() int {
	v.rw.RLock()
	defer v.rw.RUnlock()
	return len(v.v)
}

func (v *SyncVec[E]) Empty() bool {
	v.rw.RLock()
	defer v.rw.RUnlock()
	return len(v.v) == 0
}

func (v *SyncVec[E]) Cap() int {
	v.rw.RLock()
	defer v.rw.RUnlock()
	return cap(v.v)
}

func (v *SyncVec[E]) Get(index int) option.Opt[E] {
	v.rw.RLock()
	defer v.rw.RUnlock()
	if index < len(v.v) {
		return option.OptOf(v.v[index], true)
	}
	return option.OptOfEmpty[E]()
}

func (v *SyncVec[E]) First() option.Opt[E] {
	v.rw.RLock()
	defer v.rw.RUnlock()
	if len(v.v) != 0 {
		return option.OptOf(v.v[0], true)
	}
	return option.Opt[E]{}
}

func (v *SyncVec[E]) Last() option.Opt[E] {
	v.rw.RLock()
	defer v.rw.RUnlock()
	length := len(v.v)
	if length != 0 {
		return option.OptOf(v.v[length-1], true)
	}
	return option.Opt[E]{}
}

func (v *SyncVec[E]) Reverse() {
	v.rw.Lock()
	defer v.rw.Unlock()
	slices.Reverse(v.v)
}

func (v *SyncVec[E]) Clear() {
	v.rw.Lock()
	defer v.rw.Unlock()
	clear(v.v)
}

func (v *SyncVec[E]) Append(e E) {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = append(v.v, e)
}

func (v *SyncVec[E]) AppendSelf(e E) *SyncVec[E] {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = append(v.v, e)
	return v
}

func (v *SyncVec[E]) Appends(es []E) {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = append(v.v, es...)
}

func (v *SyncVec[E]) AppendsSelf(es []E) *SyncVec[E] {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = append(v.v, es...)
	return v
}

func (v *SyncVec[E]) Insert(index int, es ...E) {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = slices.Insert(v.v, index, es...)
}

func (v *SyncVec[E]) Delete(index int) {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = slices.Delete(v.v, index, index+1)
}

func (v *SyncVec[E]) DeleteOne(fn func(E) bool) {
	v.rw.Lock()
	defer v.rw.Unlock()
	for index, e := range v.v {
		fn(e)
		v.v = slices.Delete(v.v, index, index+1)
		break
	}
}

func (v *SyncVec[E]) DeleteRange(start, end int) {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = slices.Delete(v.v, start, end)
}

func (v *SyncVec[E]) Grow(n int) {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = slices.Grow(v.v, n)
}

func (v *SyncVec[E]) Clip() {
	v.rw.Lock()
	defer v.rw.Unlock()
	v.v = slices.Clip(v.v)
}

func (v *SyncVec[E]) Clone() SyncVec[E] {
	v.rw.RLock()
	defer v.rw.RUnlock()
	return SyncVec[E]{v: slices.Clone(v.v)}
}
