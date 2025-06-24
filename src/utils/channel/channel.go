package channel

import (
	"github.com/ihangsen/common/src/utils/option"
	"log"
	"time"
)

type Sender[E any] chan<- E

type Receiver[E any] <-chan E

func New[E any](cap int) (sx Sender[E], rx Receiver[E]) {
	ch := make(chan E, cap)
	return ch, ch
}

func (c Sender[E]) Send(e E) {
	c <- e
}

func (c Sender[E]) TrySend(e E) bool {
	select {
	case c <- e:
		return true
	default:
		return false
	}
}

func (c Sender[E]) SyncSend(e E) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recover from ", r)
			ok = false
		}
	}()
	c <- e
	return true
}

func (c Sender[E]) SendTimeout(e E, timeout time.Duration) bool {
	select {
	case c <- e:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (c Sender[E]) Len() int {
	return len(c)
}

func (c Sender[E]) Cap() int {
	return cap(c)
}

func (c Sender[E]) Full() bool {
	return len(c) == cap(c)
}

func (c Sender[E]) Close() {
	close(c)
}

func (c Sender[E]) AppendSelf(element E) Sender[E] {
	c <- element
	return c
}

func (c Receiver[E]) Receive() option.Opt[E] {
	e, b := <-c
	return option.OptOf(e, b)
}

func (c Receiver[E]) TryReceive() option.Opt[E] {
	select {
	case e, b := <-c:
		return option.OptOf(e, b)
	default:
		return option.Opt[E]{}
	}
}

func (c Receiver[E]) ReceiveTimeout(timeout time.Duration) option.Opt[E] {
	select {
	case e, ok := <-c:
		return option.OptOf(e, ok)
	case <-time.After(timeout):
		return option.None[E]()
	}
}

func (c Receiver[E]) ForEach(fn func(E)) {
	for e := range c {
		fn(e)
	}
}

func (c Receiver[E]) Len() int {
	return len(c)
}

func (c Receiver[E]) Cap() int {
	return cap(c)
}

func (c Receiver[E]) Empty() bool {
	return len(c) == 0
}
