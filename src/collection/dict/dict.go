package dict

import (
	"github.com/ihangsen/common/src/collection/tuple"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/function"
	"github.com/ihangsen/common/src/utils/option"
	"sort"
)

type Dict[K comparable, V any] map[K]V

func New[K comparable, V any](cap int) Dict[K, V] {
	return make(Dict[K, V], cap)
}

func Of[K comparable, V any](m map[K]V) Dict[K, V] {
	return m
}

func OfEmpty[K comparable, V any]() Dict[K, V] {
	return make(map[K]V)
}

func (d Dict[K, V]) ForEach(fn func(K, V)) {
	for k, v := range d {
		fn(k, v)
	}
}

func (d Dict[K, V]) Len() int {
	return len(d)
}

func (d Dict[K, V]) Empty() bool {
	return d.Len() == 0
}

func (d Dict[K, V]) NotEmpty() bool {
	return d.Len() > 0
}

func (d Dict[K, V]) Load(key K) option.Opt[V] {
	v, b := d[key]
	return option.OptOf(v, b)
}

func (d Dict[K, V]) Store(key K, value V) {
	d[key] = value
}

func (d Dict[K, V]) LoadOrStore(key K, value V) V {
	v, b := d[key]
	if !b {
		d[key] = value
		return value
	}
	return v
}

func (d Dict[K, V]) Delete(key K) {
	delete(d, key)
}

func (d Dict[K, V]) LoadAndDelete(key K) option.Opt[V] {
	v, b := d[key]
	if b {
		delete(d, key)
	}
	return option.OptOf(v, b)
}

func (d Dict[K, V]) ToVec() vec.Vec[tuple.T2[K, V]] {
	arr := vec.New[tuple.T2[K, V]](d.Len())
	for k, v := range d {
		arr.Append(tuple.T2Of(k, v))
	}
	return arr
}

func (d Dict[K, V]) Keys() vec.Vec[K] {
	arr := vec.New[K](d.Len())
	for k := range d {
		arr.Append(k)
	}
	return arr
}

func (d Dict[K, V]) Values() vec.Vec[V] {
	arr := vec.New[V](d.Len())
	for _, v := range d {
		arr.Append(v)
	}
	return arr
}

func (d Dict[K, V]) Clear() {
	clear(d)
}

func (d Dict[K, V]) KSort(fn func(K, K) bool) vec.Vec[tuple.T2[K, V]] {
	keys := d.Keys()
	sort.Slice(keys, func(i, j int) bool {
		return fn(keys[i], keys[j])
	})
	return function.Map(keys, func(k K) tuple.T2[K, V] {
		return tuple.T2Of(k, d[k])
	})
}
