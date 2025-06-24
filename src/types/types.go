package types

// Unit 空结构体
type Unit struct{}

func (Unit) String() string {
	return "unit"
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Comparable interface {
	Number | ~string
}

type Iterator[E any] interface {
	ForEach(func(E))
	Len() int
	Empty() bool
}
