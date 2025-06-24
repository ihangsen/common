package promise

import (
	"fmt"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/result"
	"sync"
)

type Promises[T any] struct {
	*pins[T]
}

type pins[T any] struct {
	waiter sync.WaitGroup
	m      sync.Mutex
	res    vec.Vec[result.Res[T]]
}

func All[T any](fns ...func() result.Res[T]) Promises[T] {
	length := len(fns)
	p := Promises[T]{&pins[T]{res: vec.New[result.Res[T]](length)}}
	for i := 0; i < length; i++ {
		p.res.Append(result.Res[T]{})
	}
	p.waiter.Add(length)
	for i0, fn := range fns {
		i1 := i0
		go func() {
			defer p.waiter.Done()
			defer func() {
				if r := recover(); r != nil {
					p.m.Lock()
					defer p.m.Unlock()
					p.res.Insert(i1, result.Err[T](fmt.Errorf("panic: %v", r)))
				}
			}()
			r := fn()
			p.m.Lock()
			defer p.m.Unlock()
			p.res.Insert(i1, r)
		}()
	}
	return p
}

func (p Promises[T]) Await() vec.Vec[result.Res[T]] {
	p.waiter.Wait()
	return p.res
}
