package promise

import (
	"fmt"
	"github.com/ihangsen/common/src/result"
	"sync"
)

type Promise3[T0, T1, T2 any] struct {
	*pin3[T0, T1, T2]
}

type pin3[T0, T1, T2 any] struct {
	waiter sync.WaitGroup
	res0   result.Res[T0]
	res1   result.Res[T1]
	res2   result.Res[T2]
}

func New3[T0, T1, T2 any](fn0 func() T0, fn1 func() T1, fn2 func() T2) Promise3[T0, T1, T2] {
	p := Promise3[T0, T1, T2]{&pin3[T0, T1, T2]{}}
	p.waiter.Add(3)
	go func() {
		defer p.waiter.Done()
		defer func() {
			if r := recover(); r != nil {
				p.res0 = result.Err[T0](fmt.Errorf("panic: %v", r))
			}
		}()
		p.res0 = result.Ok(fn0())
	}()
	go func() {
		defer p.waiter.Done()
		defer func() {
			if r := recover(); r != nil {
				p.res1 = result.Err[T1](fmt.Errorf("panic: %v", r))
			}
		}()
		p.res1 = result.Ok(fn1())
	}()
	go func() {
		defer p.waiter.Done()
		defer func() {
			if r := recover(); r != nil {
				p.res2 = result.Err[T2](fmt.Errorf("panic: %v", r))
			}
		}()
		p.res2 = result.Ok(fn2())
	}()
	return p
}

func (p Promise3[T0, T1, T2]) Await3() (result.Res[T0], result.Res[T1], result.Res[T2]) {
	p.waiter.Wait()
	return p.res0, p.res1, p.res2
}

func (p Promise3[T0, T1, T2]) TryAwait3() (T0, T1, T2) {
	p.waiter.Wait()
	return p.res0.Get(), p.res1.Get(), p.res2.Get()
}
