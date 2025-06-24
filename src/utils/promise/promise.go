package promise

import (
	"fmt"
	"github.com/ihangsen/common/src/result"
	"sync"
)

type Promise[T any] struct {
	*pin[T]
}

type pin[T any] struct {
	waiter sync.WaitGroup
	res    result.Res[T]
}

func New[T any](fn func() result.Res[T]) Promise[T] {
	p := Promise[T]{&pin[T]{}}
	p.waiter.Add(1)
	go func() {
		defer p.waiter.Done()
		defer func() {
			if r := recover(); r != nil {
				p.res = result.Err[T](fmt.Errorf("panic: %v", r))
			}
		}()
		p.res = fn()
	}()
	return p
}

func (p Promise[T]) Await() result.Res[T] {
	p.waiter.Wait()
	return p.res
}
