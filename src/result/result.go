package result

type Res[T any] struct {
	value T
	err   error
}

func Ok[T any](value T) Res[T] {
	return Res[T]{value: value}
}

func Err[T any](err error) Res[T] {
	return Res[T]{err: err}
}

func (r Res[T]) IsOk() bool {
	return r.err == nil
}

func (r Res[T]) IsErr() bool {
	return r.err != nil
}

func (r Res[T]) Get() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

func (r Res[T]) GetOr(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}
	return r.value
}

func (r Res[T]) Map(f func(T) T) Res[T] {
	if r.IsOk() {
		return Ok(f(r.value))
	}
	return r
}

func (r Res[T]) MapErr(f func(error) error) Res[T] {
	if r.IsErr() {
		return Err[T](f(r.err))
	}
	return r
}

func (r Res[T]) Expect() {
	if r.IsErr() {
		panic(r.err)
	}
}
