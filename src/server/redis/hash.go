package redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/dict"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/types"
	"github.com/ihangsen/common/src/utils/option"
)

func HDel(key string, fields vec.Vec[string]) {
	catch.Try(client.HDel(context.Background(), key, fields...).Err())
}

func HExists(key string, field string) bool {
	return catch.Try1(client.Do(context.Background(), "hexists", key, field).Bool())
}

func HGet[T types.Number | string](key string, field string) option.Opt[T] {
	result, err := client.Do(context.Background(), "hget", key, field).Result()
	return option.OptOf[T](result.(T), try(err))
}

func HGetObj[T any](key string, field string) option.Opt[T] {
	result, err := client.Do(context.Background(), "hget", key, field).Text()
	if !try(err) {
		return option.OptOfEmpty[T]()
	}
	t := new(T)
	catch.Try(sonic.Unmarshal([]byte(result), t))
	return option.Some[T](*t)
}

func HIncrBy(key string, field string, incr int64) int64 {
	return catch.Try1(client.Do(context.Background(), "hincrby", key, field, incr).Int64())
}

func HKeys(key string) vec.Vec[string] {
	result := catch.Try1(client.Do(context.Background(), "hkeys", key).StringSlice())
	return vec.Of[string](result...)
}

func HLen(key string) int64 {
	return catch.Try1(client.Do(context.Background(), "hlen", key).Int64())
}

func HMGet[T types.Number | string](key string, fields vec.Vec[string]) vec.Vec[T] {
	result := catch.Try1(client.HMGet(context.Background(), key, fields...).Result())
	vector := vec.New[T](len(result))
	for _, v := range result {
		vector.Append(v.(T))
	}
	return vector
}

func HMGetObj[T any](key string, fields vec.Vec[string]) vec.Vec[T] {
	result := catch.Try1(client.HMGet(context.Background(), key, fields...).Result())
	vector := vec.New[T](2)
	for _, v := range result {
		t := new(T)
		if v != nil {
			catch.Try(sonic.Unmarshal([]byte(v.(string)), t))
		}
		vector.Append(*t)
	}
	return vector
}

func HMGetDict[T types.Number | string](key string, fields vec.Vec[string]) dict.Dict[string, T] {
	result := catch.Try1(client.HMGet(context.Background(), key, fields...).Result())
	dictResult := dict.New[string, T](len(result))
	for i := range fields {
		dictResult.Store(fields[i], result[i].(T))
	}
	return dictResult
}

func HMGetDictObj[T any](key string, fields vec.Vec[string]) dict.Dict[string, T] {
	result := catch.Try1(client.HMGet(context.Background(), key, fields...).Result())
	dictResult := dict.New[string, T](len(result))
	for i, v := range result {
		t := new(T)
		if v != nil {
			catch.Try(sonic.Unmarshal([]byte(v.(string)), t))
		}
		dictResult.Store(fields[i], *t)
	}
	return dictResult
}

func HSetObj[T any](key string, field string, value T) {
	_value := catch.Try1(sonic.Marshal(value))
	catch.Try(client.Do(context.Background(), "hset", key, field, _value).Err())
}

func HSet[T types.Number | string](key string, field string, value T) {
	catch.Try(client.Do(context.Background(), "hset", key, field, value).Err())
}

// HSetNXObj 只有在字段 field 不存在时，设置哈希表字段的值。
func HSetNXObj[T any](key string, field string, value T) {
	_value := catch.Try1(sonic.Marshal(value))
	catch.Try(client.Do(context.Background(), "hsetnx", key, field, _value).Err())
}

func HSetNX[T types.Number | string](key string, field string, value T) {
	catch.Try(client.Do(context.Background(), "hsetnx", key, field, value).Err())
}
