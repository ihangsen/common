package bit_map

import (
	"github.com/ihangsen/common/src/collection/set"
	"github.com/ihangsen/common/src/collection/vec"
	. "github.com/ihangsen/common/src/types"
	"unsafe"
)

type BitMap[T Integer] struct {
	value T
}

func BitMapNew[T Integer](value T) BitMap[T] {
	return BitMap[T]{value}
}

func (bm *BitMap[T]) Value() T {
	return bm.value
}

// Set 把T的第index位置为b
func (bm *BitMap[T]) Set(index int, b bool) {
	v := bm.value
	if b {
		v |= T(1) << index
	} else {
		v &= ^(T(1) << index)
	}
	bm.value = v
}

// Get 获取T中的第index位置的值
func (bm *BitMap[T]) Get(index int) bool {
	if index >= int(unsafe.Sizeof(bm.value))*8 {
		return false
	}
	return bm.get(index)
}

func (bm *BitMap[T]) Count() int {
	count := 0
	len_ := int(unsafe.Sizeof(bm.value)) * 8
	for i := range len_ {
		if bm.get(i) {
			count += 1
		}
	}
	return count
}

func (bm *BitMap[T]) get(index int) bool {
	return (bm.value>>index)&1 == 1
}

type BytesBitMap struct {
	value []byte
}

func BytesBitMapNew(value []byte) BytesBitMap {
	return BytesBitMap{value}
}

func (bm *BytesBitMap) Value() []byte {
	return bm.value
}

func (bm *BytesBitMap) String() string {
	return string(bm.value)
}

// Set 把T的第index位置为b
func (bm *BytesBitMap) Set(index int, b bool) {
	outIndex := index / 8
	if outIndex >= len(bm.value) {
		value := make([]byte, outIndex+1)
		copy(value, bm.value)
		bm.value = value
	}
	v := bm.value[outIndex]
	inIndex := index % 8
	if b {
		v |= byte(1) << inIndex
	} else {
		v &= ^(byte(1) << inIndex)
	}
	bm.value[outIndex] = v
}

// Get 获取T中的第index位置的值
func (bm *BytesBitMap) Get(index int) bool {
	if index >= len(bm.value)*8 {
		return false
	}
	return bm.get(index)
}

func (bm *BytesBitMap) Count() int {
	count := 0
	length := len(bm.value) * 8
	for i := range length {
		if bm.get(i) {
			count += 1
		}
	}
	return count
}

func (bm *BytesBitMap) ToSet() set.Set[uint16] {
	length := len(bm.value) * 8
	res := set.New[uint16](length / 2)
	for i := range length {
		if bm.get(i) {
			res.Insert(uint16(i))
		}
	}
	return res
}

func (bm *BytesBitMap) ToVec() vec.Vec[uint16] {
	length := len(bm.value) * 8
	res := vec.New[uint16](length / 2)
	for i := range length {
		if bm.get(i) {
			res.Append(uint16(i))
		}
	}
	return res
}

func (bm *BytesBitMap) Len() int {
	return len(bm.value) * 8
}

func (bm *BytesBitMap) get(index int) bool {
	return (bm.value[index/8]>>(index%8))&1 == 1
}
