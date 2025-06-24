package redis

import (
	"context"
	"github.com/ihangsen/common/src/catch"
)

func SetBit(key string, offset int64, value int) int64 {
	return catch.Try1(client.SetBit(context.Background(), key, offset, value).Result())
}

func GetBit(key string, offset int64) int64 {
	return catch.Try1(client.GetBit(context.Background(), key, offset).Result())
}
