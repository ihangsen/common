package redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/types"
)

func LPush[T string | types.Number](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "LPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = value
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func LPushObj[T any](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "LPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = catch.Try1(sonic.Marshal(value))
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func RPush[T string | types.Number](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "RPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = value
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func RPushObj[T any](key string, values vec.Vec[T]) {
	args := make([]any, 2+len(values))
	args[0] = "RPUSH"
	args[1] = key
	for i, value := range values {
		args[i+2] = catch.Try1(sonic.Marshal(value))
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func LRange[T string | types.Number](key string, start, stop int64) vec.Vec[T] {
	result := catch.Try1(client.Do(context.Background(), "LRANGE", key, start, stop).Slice())
	temp := vec.New[T](len(result))
	for _, v := range result {
		temp.Append(v.(T))
	}
	return temp
}

func LRangeObj[T any](key string, start, stop int64) vec.Vec[T] {
	result := catch.Try1(client.Do(context.Background(), "LRANGE", key, start, stop).Slice())
	temp := vec.New[T](len(result))
	for _, v := range result {
		t := new(T)
		catch.Try(sonic.Unmarshal([]byte(v.(string)), t))
		temp.Append(*t)
	}
	return temp
}

func LRem[T string | types.Number](key string, value T) {
	catch.Try(client.Do(context.Background(), "LREM", key, 0, value).Err())
}

func LRemObj[T any](key string, value T) {
	catch.Try(client.Do(context.Background(), "LREM", key, 0, catch.Try1(sonic.Marshal(value))).Err())
}

func LLen(key string) int {
	return catch.Try1(client.Do(context.Background(), "LLEN", key).Int())
}
