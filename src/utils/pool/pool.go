package pool

import "sync"

type Pool[T any] struct {
	p sync.Pool
}

func New[T any](fn func() T) Pool[T] {
	return Pool[T]{p: sync.Pool{
		New: func() any { return fn() },
	}}
}

func (p *Pool[T]) Put(t T) {
	p.p.Put(t)
}

func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}
