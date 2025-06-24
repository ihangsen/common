package vec

import (
	"github.com/ihangsen/common/src/utils/option"
	"slices"
)

type Vec[E any] []E

func New[E any](cap int) Vec[E] {
	return make(Vec[E], 0, cap)
}

func Of[E any](arr ...E) Vec[E] {
	return arr
}

func OfEmpty[E any]() Vec[E] {
	return make(Vec[E], 0)
}

func (v Vec[E]) ForEach(fn func(E)) {
	for _, e := range v {
		fn(e)
	}
}

func (v Vec[E]) Len() int {
	return len(v)
}

func (v Vec[E]) Empty() bool {
	return len(v) == 0
}

func (v Vec[E]) Cap() int {
	return cap(v)
}

func (v Vec[E]) Get(index int) option.Opt[E] {
	if index < v.Len() {
		return option.OptOf(v[index], true)
	}
	return option.OptOfEmpty[E]()
}

func (v Vec[E]) First() option.Opt[E] {
	if !v.Empty() {
		return option.OptOf(v[0], true)
	}
	return option.Opt[E]{}
}

func (v Vec[E]) Last() option.Opt[E] {
	if !v.Empty() {
		return option.OptOf(v[len(v)-1], true)
	}
	return option.Opt[E]{}
}

func (v Vec[E]) Reverse() {
	slices.Reverse(v)
}

func (v Vec[E]) Clear() {
	clear(v)
}

func (v *Vec[E]) Append(element E) {
	*v = append(*v, element)
}

func (v Vec[E]) AppendSelf(element E) Vec[E] {
	return append(v, element)
}

func (v *Vec[E]) Appends(elements Vec[E]) {
	*v = append(*v, elements...)
}

func (v Vec[E]) AppendsSelf(elements Vec[E]) Vec[E] {
	return append(v, elements...)
}

func (v *Vec[E]) Insert(index int, elements ...E) {
	*v = slices.Insert(*v, index, elements...)
}

func (v *Vec[E]) Delete(index int) {
	*v = slices.Delete(*v, index, index+1)
}

func (v *Vec[E]) DeleteRange(start, end int) {
	*v = slices.Delete(*v, start, end)
}

func (v *Vec[E]) Grow(n int) {
	*v = slices.Grow(*v, n)
}

func (v *Vec[E]) Clip() {
	*v = slices.Clip(*v)
}

func (v Vec[E]) Clone() Vec[E] {
	return slices.Clone(v)
}
