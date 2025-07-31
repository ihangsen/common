package redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/function"
	"github.com/ihangsen/common/src/types"
)

func LPush[T string | types.Number](key string, values vec.Vec[T]) {
	catch.Try(client.LPush(context.Background(), key, values).Err())
}

func LPushObj[T any](key string, values vec.Vec[T]) {
	catch.Try(client.LPush(context.Background(), key, function.Map(values, func(v T) string {
		return catch.Try1(sonic.MarshalString(v))
	})).Err())
}

func RPush[T string | types.Number](key string, values vec.Vec[T]) {
	catch.Try(client.RPush(context.Background(), key, values).Err())
}

func RPushObj[T any](key string, values vec.Vec[T]) {
	catch.Try(client.RPush(context.Background(), key, function.Map(values, func(v T) string {
		return catch.Try1(sonic.MarshalString(v))
	})).Err())
}

func LRange[T string | types.Number](key string, start, stop int64) vec.Vec[T] {
	result := catch.Try1(client.Do(context.Background(), key, start, stop).Slice())
	return function.Map(result, func(t any) T {
		return t.(T)
	})
}

func LRangeObj[T any](key string, start, stop int64) vec.Vec[T] {
	result := catch.Try1(client.Do(context.Background(), key, start, stop).Slice())
	temp := vec.New[T](len(result))
	for i, v := range result {
		t := new(T)
		catch.Try(sonic.Unmarshal(v.([]byte), t))
		temp[i] = *t
	}
	return temp
}

func LRem(key string, count int64, value any) {
	catch.Try(client.LRem(context.Background(), key, count, value).Err())
}
