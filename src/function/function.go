package function

import (
	"github.com/ihangsen/common/src/collection/dict"
	"github.com/ihangsen/common/src/collection/set"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/types"
)

func Map[T, R any, TS ~[]T](ts TS, fn func(T) R) vec.Vec[R] {
	rs := vec.New[R](len(ts))
	for _, t := range ts {
		rs.Append(fn(t))
	}
	return rs
}

func MapIndex[T, R any, TS ~[]T](ts TS, fn func(int, T) R) vec.Vec[R] {
	rs := vec.New[R](len(ts))
	for i, t := range ts {
		rs.Append(fn(i, t))
	}
	return rs
}

func ToSet[T any, K comparable, TS ~[]T](ts TS, fn func(T) K) set.Set[K] {
	s := set.New[K](len(ts))
	for _, t := range ts {
		s.Insert(fn(t))
	}
	return s
}

func ToMap[K comparable, V any, TS types.Iterator[V]](ts TS, kFn func(V) K) dict.Dict[K, V] {
	d := dict.New[K, V](ts.Len())
	ts.ForEach(func(t V) {
		d.Store(kFn(t), t)
	})
	return d
}

func ToMapKV[K comparable, V, T any, TS types.Iterator[T]](ts TS, kvFn func(T) (K, V)) dict.Dict[K, V] {
	d := dict.New[K, V](ts.Len())
	ts.ForEach(func(t T) {
		d.Store(kvFn(t))
	})
	return d
}

func SyncToMapKV[K comparable, V, T any, TS types.Iterator[T]](ts TS, kvFn func(T) (K, V)) *dict.SyncDict[K, V] {
	d := dict.SyncDict[K, V]{}
	ts.ForEach(func(t T) {
		d.Store(kvFn(t))
	})
	return &d
}

func GroupBy[K comparable, V any, TS types.Iterator[V]](ts TS, kFn func(V) K) dict.Dict[K, vec.Vec[V]] {
	d := dict.New[K, vec.Vec[V]](ts.Len())
	ts.ForEach(func(t V) {
		d.Store(kFn(t), append(d[kFn(t)], t))
	})
	return d
}

func GroupByKV[K comparable, V, T any, TS types.Iterator[T]](ts TS, kvFn func(T) (K, V)) dict.Dict[K, vec.Vec[V]] {
	d := dict.New[K, vec.Vec[V]](ts.Len())
	ts.ForEach(func(t T) {
		k, v := kvFn(t)
		d.Store(k, append(d[k], v))
	})
	return d
}

func Reduce[T, R any, TS types.Iterator[T]](ts TS, seed R, fn func(R, T) R) R {
	ts.ForEach(func(t T) {
		seed = fn(seed, t)
	})
	return seed
}

func Filter[T any, TS types.Iterator[T]](ts TS, fn func(T) bool) vec.Vec[T] {
	rs := vec.New[T](filterLen(ts.Len()))
	ts.ForEach(func(t T) {
		if fn(t) {
			rs.Append(t)
		}
	})
	return rs
}

func filterLen(len_ int) int {
	switch {
	case len_ < 8:
		return len_
	case len_ < 32:
		len2 := len_ / 2
		if len2*2 < len_ {
			len2 += 1
		}
		return len2
	default:
		len4 := len_ / 4
		if len4*4 < len_ {
			len4 += 1
		}
		return len4
	}
}
