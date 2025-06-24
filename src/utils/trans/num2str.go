package trans

import (
	"strconv"
)

func U82Str(num uint8) string {
	return strconv.FormatUint(uint64(num), 10)
}

func I82Str(num int8) string {
	return strconv.FormatInt(int64(num), 10)
}

func U162Str(num uint16) string {
	return strconv.FormatUint(uint64(num), 10)
}

func I162Str(num int16) string {
	return strconv.FormatInt(int64(num), 10)
}

func U322Str(num uint32) string {
	return strconv.FormatUint(uint64(num), 10)
}

func UInt2Str(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func I322Str(num int32) string {
	return strconv.FormatInt(int64(num), 10)
}

func Int2Str(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func U642Str(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func I642Str(num int64) string {
	return strconv.FormatInt(num, 10)
}

func F322Str(num float32) string {
	return strconv.FormatFloat(float64(num), 'g', -1, 32)
}

func F642Str(num float64) string {
	return strconv.FormatFloat(num, 'g', -1, 64)
}
