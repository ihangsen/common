package promise

import (
	"fmt"
	"github.com/ihangsen/common/src/result"
	"sync"
)

type Promise5[T0, T1, T2, T3, T4 any] struct {
	*pin5[T0, T1, T2, T3, T4]
}

type pin5[T0, T1, T2, T3, T4 any] struct {
	waiter sync.WaitGroup
	res0   result.Res[T0]
	res1   result.Res[T1]
	res2   result.Res[T2]
	res3   result.Res[T3]
	res4   result.Res[T4]
}

func New5[T0, T1, T2, T3, T4 any](fn0 func() T0, fn1 func() T1, fn2 func() T2, fn3 func() T3, fn4 func() T4) Promise5[T0, T1, T2, T3, T4] {
	p := Promise5[T0, T1, T2, T3, T4]{&pin5[T0, T1, T2, T3, T4]{}}
	p.waiter.Add(5)
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
	go func() {
		defer p.waiter.Done()
		defer func() {
			if r := recover(); r != nil {
				p.res3 = result.Err[T3](fmt.Errorf("panic: %v", r))
			}
		}()
		p.res3 = result.Ok(fn3())
	}()
	go func() {
		defer p.waiter.Done()
		defer func() {
			if r := recover(); r != nil {
				p.res3 = result.Err[T3](fmt.Errorf("panic: %v", r))
			}
		}()
		p.res4 = result.Ok(fn4())
	}()
	return p
}

func (p Promise5[T0, T1, T2, T3, T4]) Await5() (result.Res[T0], result.Res[T1], result.Res[T2], result.Res[T3], result.Res[T4]) {
	p.waiter.Wait()
	return p.res0, p.res1, p.res2, p.res3, p.res4
}

func (p Promise5[T0, T1, T2, T3, T4]) TryAwait5() (T0, T1, T2, T3, T4) {
	p.waiter.Wait()
	return p.res0.Get(), p.res1.Get(), p.res2.Get(), p.res3.Get(), p.res4.Get()
}
