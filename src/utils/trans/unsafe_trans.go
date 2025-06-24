package trans

import (
	"encoding/binary"
	. "github.com/ihangsen/common/src/types"
	"unsafe"
)

// UnsafeTrans 对值类型做二进制重新解释 谨慎使用
func UnsafeTrans[R, T any](t T) R {
	return *(*R)(unsafe.Pointer(&t))
}

// UnsafeBytes2Str 为了减少内存copy 使用的不安全转换 谨慎使用 传入的bytes不允许再做修改
func UnsafeBytes2Str(bytes []byte) string {
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

// UnsafeStr2Bytes 为了减少内存copy 使用的不安全转换 谨慎使用
func UnsafeStr2Bytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

// UnsafeBytes2NumberBe 二进制转数字 大端
func UnsafeBytes2NumberBe[T Number](bytes []byte) T {
	t := *new(T)
	size := unsafe.Sizeof(t)
	switch size {
	case 1:
		t = UnsafeTrans[T](bytes[0])
	case 2:
		t = UnsafeTrans[T](binary.BigEndian.Uint16(bytes))
	case 4:
		t = UnsafeTrans[T](binary.BigEndian.Uint32(bytes))
	case 8:
		t = UnsafeTrans[T](binary.BigEndian.Uint64(bytes))
	}
	return t
}

// UnsafeBytes2NumberLe 二进制转数字 小端
func UnsafeBytes2NumberLe[T Number](bytes []byte) T {
	t := *new(T)
	size := unsafe.Sizeof(t)
	switch size {
	case 1:
		t = UnsafeTrans[T](bytes[0])
	case 2:
		t = UnsafeTrans[T](binary.LittleEndian.Uint16(bytes))
	case 4:
		t = UnsafeTrans[T](binary.LittleEndian.Uint32(bytes))
	case 8:
		t = UnsafeTrans[T](binary.LittleEndian.Uint64(bytes))
	}
	return t
}
