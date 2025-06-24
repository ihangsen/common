package promise

import (
	"fmt"
	"github.com/ihangsen/common/src/result"
	"sync"
)

type Promise2[T0, T1 any] struct {
	*pin2[T0, T1]
}

type pin2[T0, T1 any] struct {
	waiter sync.WaitGroup
	res0   result.Res[T0]
	res1   result.Res[T1]
}

func New2[T0, T1 any](fn0 func() T0, fn1 func() T1) Promise2[T0, T1] {
	p := Promise2[T0, T1]{&pin2[T0, T1]{}}
	p.waiter.Add(2)
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
	return p
}

func (p Promise2[T0, T1]) Await2() (result.Res[T0], result.Res[T1]) {
	p.waiter.Wait()
	return p.res0, p.res1
}

func (p Promise2[T0, T1]) TryAwait2() (T0, T1) {
	p.waiter.Wait()
	return p.res0.Get(), p.res1.Get()
}
