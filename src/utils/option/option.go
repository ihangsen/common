package option

import (
	"errors"
	"fmt"
	"github.com/ihangsen/common/src/res"
	"github.com/ihangsen/common/src/status"
)

type Opt[T any] struct {
	V T
	B bool
}

func OptOf[T any](t T, b bool) Opt[T] {
	return Opt[T]{t, b}
}

func OptOfEmpty[T any]() Opt[T] {
	return Opt[T]{}
}

func Some[T any](t T) Opt[T] {
	return Opt[T]{t, true}
}

func None[T any]() Opt[T] {
	return Opt[T]{}
}

func (o Opt[T]) Unravel() (T, bool) {
	return o.V, o.B
}

func (o Opt[T]) IsSome() bool {
	return o.B
}

func (o Opt[T]) IsNone() bool {
	return !o.B
}

func (o Opt[T]) Expect(msg string) T {
	if o.IsSome() {
		return o.V
	}
	panic(res.Err{Code: status.CustomCode, Msg: msg})
}

func (o Opt[T]) ExpectCode(code uint8) T {
	if o.IsSome() {
		return o.V
	}
	panic(res.Err{Code: code})
}

func (o Opt[T]) ExpectErr(err res.Err) T {
	if o.IsSome() {
		return o.V
	}
	panic(err)
}

func (o Opt[T]) Get() T {
	if o.IsSome() {
		return o.V
	}
	panic(errors.New("option is none"))
}

func (o Opt[T]) GetOr(t T) T {
	if o.IsSome() {
		return o.V
	}
	return t
}

func (o Opt[T]) GetOrDefault() T {
	if o.IsSome() {
		return o.V
	}
	return *new(T)
}

func (o Opt[T]) GetElse(fn func() T) T {
	if o.IsSome() {
		return o.V
	}
	return fn()
}

func (o Opt[T]) Map(fn func(t T)) {
	if o.IsSome() {
		fn(o.V)
	}
}

func (o Opt[T]) MapOr(t T, fn func(t T) T) T {
	if o.IsSome() {
		return fn(o.V)
	} else {
		return t
	}
}

func (o Opt[T]) MapOrElse(fn0 func(), fn1 func(t T)) {
	if o.IsSome() {
		fn1(o.V)
	} else {
		fn0()
	}
}

func (o Opt[T]) Or(fn func() Opt[T]) Opt[T] {
	if o.IsSome() {
		return fn()
	}
	return o
}

func (o Opt[T]) Else(fn func(t T) Opt[T]) Opt[T] {
	if o.IsSome() {
		return o
	}
	return fn(o.V)
}

func (o Opt[T]) OrElse(fn0 func() Opt[T], fn1 func(t T) Opt[T]) Opt[T] {
	if o.IsSome() {
		return fn0()
	}
	return fn1(o.V)
}

func (o Opt[T]) String() string {
	if o.IsSome() {
		return fmt.Sprintf("some(%v)", o.V)
	}
	return "none"
}

// NzOpt Non zero option
type NzOpt[T comparable] struct {
	V T
}

func NzOptOf[T comparable](t T) NzOpt[T] {
	return NzOpt[T]{t}
}

func NzOptOfEmpty[T comparable]() NzOpt[T] {
	return NzOpt[T]{}
}

func (o NzOpt[T]) D() (T, bool) {
	return o.V, o.V != *new(T)
}

func (o NzOpt[T]) IsSome() bool {
	return o.V != *new(T)
}

func (o NzOpt[T]) IsNone() bool {
	return o.V == *new(T)
}

func (o NzOpt[T]) Expect(msg string) T {
	if o.IsSome() {
		return o.V
	}
	panic(res.Err{Code: status.CustomCode, Msg: msg})
}

func (o NzOpt[T]) ToOpt() Opt[T] {
	if o.IsSome() {
		return Some(o.V)
	}
	return None[T]()
}

// Get 获取值 如果为none 则会panic
func (o NzOpt[T]) Get() T {
	if o.IsSome() {
		return o.V
	}
	panic(errors.New("option is none"))
}

func (o NzOpt[T]) GetOr(t T) T {
	if o.IsSome() {
		return o.V
	}
	return t
}

func (o NzOpt[T]) GetElse(fn func() T) T {
	if o.IsSome() {
		return o.V
	}
	return fn()
}

func (o NzOpt[T]) MapOrElse(fn0 func(), fn1 func(t T)) {
	if o.IsSome() {
		fn1(o.V)
	} else {
		fn0()
	}
}

func (o NzOpt[T]) Map(fn func(t T)) {
	if o.IsSome() {
		fn(o.V)
	}
}

func (o NzOpt[T]) String() string {
	if o.IsSome() {
		return fmt.Sprintf("some(%v)", o.V)
	}
	return "none"
}
