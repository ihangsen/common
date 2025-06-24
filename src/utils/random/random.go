package random

import (
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/ihangsen/common/src/catch"
	trans2 "github.com/ihangsen/common/src/utils/trans"
	"math/rand/v2"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	pcg    = rand.New(rand.NewPCG(uint64(time.Now().UnixMilli()), uint64(time.Now().UnixMilli())))
	length = len(charset)
)

func RandomNum6() string {
	return trans2.Int2Str(pcg.IntN(900000) + 100000)
}
func RandNumEndI64(end int64) int64 {
	return pcg.Int64N(end)
}

func RandomI64(start, count int64) int64 {
	return pcg.Int64N(count) + start
}

func UUID() string {
	u := catch.Try1(uuid.NewUUID())
	result := make([]byte, 32)
	hex.Encode(result, u[:])
	return trans2.UnsafeBytes2Str(result)
}

func str(l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = charset[pcg.IntN(length)]
	}
	return trans2.UnsafeBytes2Str(b)
}
