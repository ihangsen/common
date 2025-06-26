package redis

import (
	"context"
	"errors"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/log"
	"github.com/ihangsen/common/src/utils/option"
	"github.com/ihangsen/common/src/utils/trans"
	"github.com/redis/go-redis/v9"
	"time"
)

// XAck 确认消息
func XAck(key, group string, ids vec.Vec[string]) {
	args := make([]any, 0, 3+len(ids))
	args = append(args, "xack", key, group)
	for _, id := range ids {
		args = append(args, id)
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

// XClaim 重新将消息转给指定消费者处理,只返回ID
func XClaim(key, group, consumer, id string, minIdleTime int64) {
	catch.Try(client.Do(context.Background(), "xclaim", key, group, consumer, minIdleTime, id).Err())
}

func XAutoClaim(key, group, consumer string, minIdleTime int64) {
	catch.Try(client.Do(context.Background(), "xautoclaim", key, group, consumer, minIdleTime, "0-0", "count", 10).Err())
}

func XAdd[T any](key string, t T, maxLen int64) string {
	m := trans.Obj2Json(t)
	args := make([]any, 0, 7)
	args = append(args, "xadd", key, "maxlen", maxLen, "*")
	args = append(args, "data", m)
	result := catch.Try1(client.Do(context.Background(), args...).Result())
	return result.(string)
}

func XDel(key string, ids vec.Vec[string]) {
	args := make([]any, 0, 2+len(ids))
	args = append(args, "xdel", key)
	for _, id := range ids {
		args = append(args, id)
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

func XGroupCreate(key, group string) {
	catch.Try(client.Do(context.Background(), "xgroup", "create", key, group, "$", "MKSTREAM").Err())
}

type XMessage[T any] struct {
	ID  string
	Msg T
}

func XReadGroup[T any](key, group, consumer string) option.Opt[XMessage[T]] {
	args := &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumer,
		Streams:  []string{key, ">"}, // 处理最先进来的消息
		Count:    1,
		Block:    time.Duration(500) * time.Millisecond,
		NoAck:    true,
	}
	xStreams, err := client.XReadGroup(context.Background(), args).Result()
	if err != nil || len(xStreams) <= 0 {
		if errors.Is(err, redis.Nil) {
			return option.OptOfEmpty[XMessage[T]]()
		}
		log.Zap.Errorf("获取消息失败: key:%s,group:%s,err:%s", key, group, err.Error())
		return option.OptOfEmpty[XMessage[T]]()
	}
	stream := xStreams[0]
	messages := stream.Messages[0]
	t := trans.Map2Obj[T](messages.Values)
	return option.Some(XMessage[T]{ID: messages.ID, Msg: t})
}

func XReadGroups[T any](key, group, consumer string, count int64) option.Opt[vec.Vec[XMessage[T]]] {
	args := &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumer,
		Streams:  []string{key, ">"}, // 处理最先进来的消息
		Count:    count,
		Block:    time.Duration(500) * time.Millisecond,
		NoAck:    true,
	}
	xStreams, err := client.XReadGroup(context.Background(), args).Result()
	if err != nil || len(xStreams) <= 0 {
		if errors.Is(err, redis.Nil) {
			return option.OptOfEmpty[vec.Vec[XMessage[T]]]()
		}
		log.Zap.Errorf("获取消息失败: key:%s,group:%s,err:%s", key, group, err.Error())
		return option.OptOfEmpty[vec.Vec[XMessage[T]]]()
	}
	stream := xStreams[0]
	messages := stream.Messages
	result := make(vec.Vec[XMessage[T]], len(messages))
	for i, message := range messages {
		result[i] = XMessage[T]{ID: message.ID, Msg: trans.Json2Obj[T](message.Values["data"].(string))}
	}
	return option.Some(result)
}

func XExists(key string, groupKey string) bool {
	result, err := client.XInfoGroups(context.Background(), key).Result()
	if err != nil {
		return false
	}
	for _, group := range result {
		if group.Name == groupKey {
			return true
		}
	}
	return false
}
