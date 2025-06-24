package redis

import (
	"context"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/utils/option"
	"github.com/ihangsen/common/src/utils/trans"
)

func SetPx[T any](key string, value T, expire int64) {
	catch.Try(client.Do(context.Background(), "set", key, value, "px", expire).Err())
}

func Set[T any](key string, value T) {
	catch.Try(client.Do(context.Background(), "set", key, value).Err())
}

func GetI64(key string) option.Opt[int64] {
	result, err := client.Get(context.Background(), key).Int64()
	return option.OptOf[int64](result, try(err))
}

func GetU32(key string) option.Opt[uint32] {
	result, err := client.Get(context.Background(), key).Int()
	return option.OptOf[uint32](uint32(result), try(err))
}

func GetU8(key string) option.Opt[uint8] {
	result, err := client.Get(context.Background(), key).Int()
	return option.OptOf[uint8](uint8(result), try(err))
}

func GetString(key string) option.Opt[string] {
	result, err := client.Get(context.Background(), key).Result()
	return option.OptOf[string](result, try(err))
}

func IncrBy(key string, increment uint32) int64 {
	return catch.Try1(client.IncrBy(context.Background(), key, int64(increment)).Result())
}

func MGet2(key0, key1 string) (option.NzOpt[int64], option.NzOpt[string]) {
	arr, err := client.MGet(context.Background(), key0, key1).Result()
	if try(err) {
		t0 := option.NzOptOfEmpty[int64]()
		if arr[0] != nil {
			t0 = option.NzOptOf[int64](trans.Str2I64(arr[0].(string)))
		}
		t1 := option.NzOptOfEmpty[string]()
		if arr[1] != nil {
			t1 = option.NzOptOf[string](arr[1].(string))
		}
		return t0, t1
	}
	return option.NzOptOfEmpty[int64](), option.NzOptOfEmpty[string]()
}
