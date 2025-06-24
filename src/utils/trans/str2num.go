package trans

import (
	"github.com/ihangsen/common/src/catch"
	"strconv"
)

func Str2U8(str string) uint8 {
	return uint8(catch.Try1(strconv.ParseUint(str, 10, 8)))
}

func Str2I8(str string) int8 {
	return int8(catch.Try1(strconv.ParseInt(str, 10, 8)))
}

func Str2U16(str string) uint16 {
	return uint16(catch.Try1(strconv.ParseUint(str, 10, 16)))
}

func Str2I16(str string) int16 {
	return int16(catch.Try1(strconv.ParseInt(str, 10, 16)))
}

func Str2U32(str string) uint32 {
	return uint32(catch.Try1(strconv.ParseUint(str, 10, 32)))
}

func Str2I32(str string) int32 {
	return int32(catch.Try1(strconv.ParseInt(str, 10, 32)))
}

func Str2U64(str string) uint64 {
	return catch.Try1(strconv.ParseUint(str, 10, 64))
}

func Str2I64(str string) int64 {
	return catch.Try1(strconv.ParseInt(str, 10, 64))
}

func Str2F32(str string) float32 {
	return float32(catch.Try1(strconv.ParseFloat(str, 32)))
}

func Str2F64(str string) float64 {
	return catch.Try1(strconv.ParseFloat(str, 64))
}
