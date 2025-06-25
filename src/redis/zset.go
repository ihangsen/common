package redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/types"
	"github.com/ihangsen/common/src/utils/option"
	"github.com/ihangsen/common/src/utils/trans"
)

type ZElement[T any] struct {
	Score  float64
	Member T
}

// ZAdd 添加元素
func ZAdd[T types.Number | string](key string, elements ...ZElement[T]) {
	args := make([]any, 0, 2+len(elements)*2)
	args = append(args, "zadd", key)
	for _, element := range elements {
		args = append(args, element.Score, element.Member)
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

// ZAddObj 添加结构体元素
func ZAddObj[T any](key string, elements vec.Vec[ZElement[T]]) {
	args := make([]any, 0, 2+len(elements)*2)
	args = append(args, "zadd", key)
	for _, element := range elements {
		member := catch.Try1(sonic.Marshal(element.Member))
		args = append(args, element.Score, member)
	}
	catch.Try(client.Do(context.Background(), args...).Err())
}

// ZCard 获取元素数量
func ZCard(key string) int64 {
	return catch.Try1(client.Do(context.Background(), "zcard", key).Int64())
}

// ZCount 获取指定范围内的元素数量
// min,max小于0为上下边界值
func ZCount(key string, min, max int64) int64 {
	tempMin := trans.I642Str(min)
	tempMax := trans.I642Str(max)
	if min < 0 {
		tempMin = "-inf"
	}
	if max < 0 {
		tempMax = "+inf"
	}
	return catch.Try1(client.Do(context.Background(), "zcount", key, tempMin, tempMax).Int64())
}

// ZDiff 差异的结果
func ZDiff(keys vec.Vec[string]) vec.Vec[string] {
	args := make([]any, 0, 2+len(keys))
	args = append(args, "zdiff", len(keys))
	for _, key := range keys {
		args = append(args, key)
	}
	result, err := client.Do(context.Background(), args...).StringSlice()
	if !try(err) {
		return vec.OfEmpty[string]()
	}
	vector := vec.New[string](len(result))
	for _, item := range result {
		vector.Append(item)
	}
	return vector
}

// ZIncrBy 增加元素的分数
func ZIncrBy[T any](key string, increment float64, member T) float64 {
	return catch.Try1(client.Do(context.Background(), "zincrby", key, increment, member).Float64())
}

// ZMScore 批量获取元素的分数
func ZMScore(key string, members vec.Vec[string]) vec.Vec[float64] {
	args := make([]any, 0, 2+len(members))
	args = append(args, "zmscore", key)
	for _, member := range members {
		args = append(args, member)
	}
	result, err := client.Do(context.Background(), args...).Slice()
	if !try(err) {
		return vec.OfEmpty[float64]()
	}
	vector := vec.New[float64](len(result))
	for _, item := range result {
		if item == nil {
			vector.Append(0)
			continue
		}
		vector.Append(item.(float64))
	}
	return vector
}

// ZRank 获取元素的排名
func ZRank(key, member string) int64 {
	result := catch.Try1(client.Do(context.Background(), "zrank", key, member).Int64())
	return result
}

// ZRem 移除元素
func ZRem[T types.Number | string](key string, members vec.Vec[T]) {
	catch.Try(client.ZRem(context.Background(), key, members).Err())
}

func ZRemObj[T any](key string, members vec.Vec[T]) {
	members0 := make([]string, len(members))
	for _, member := range members {
		members0 = append(members0, string(catch.Try1(sonic.Marshal(member))))
	}
	catch.Try(client.ZRem(context.Background(), key, members0).Err())
}

// ZRemRangeByRank 删除介于两者之间得元素
func ZRemRangeByRank(key string, start, stop int64) {
	catch.Try(client.ZRemRangeByRank(context.Background(), key, start, stop).Err())
}

// ZScore 获取元素的分数
func ZScore[T types.Number | string](key string, member T) option.Opt[float64] {
	result, err := client.Do(context.Background(), "zscore", key, member).Result()
	if try(err) {
		return option.None[float64]()
	}
	return option.Some(result.(float64))
}

// ZRevRank 获取元素的逆序排名
func ZRevRank(key, member string) int64 {
	return catch.Try1(client.Do(context.Background(), "zrevrank", key, member).Int64())
}

func zRangeObj[T any](key string, start, end int64, rev bool) vec.Vec[T] {
	args := make([]any, 0, 6)
	if start < 0 {
		args = append(args, "zrange", key, "-inf", end)
	} else {
		args = append(args, "zrange", key, start, end)
	}
	if rev {
		args = append(args, "rev")
	}
	args = append(args, "byscore")
	result := catch.Try1(client.Do(context.Background(), args...).Slice())
	vector := vec.New[T](len(result))
	for _, item := range result {
		t := new(T)
		if item != "" {
			catch.Try(sonic.Unmarshal([]byte(item.(string)), t))
		}
		vector.Append(*t)
	}
	return vector
}

func zRangeByScoreObj[T any](key string, start, end int64, rev bool) vec.Vec[ZElement[T]] {
	args := make([]any, 0, 7)
	if start < 0 {
		args = append(args, "zrange", key, "-inf", end)
	} else {
		args = append(args, "zrange", key, start, end)
	}
	if rev {
		args = append(args, "rev")
	}
	args = append(args, "withscores", "byscore")
	result := catch.Try1(client.Do(context.Background(), args...).Slice())
	vector := vec.New[ZElement[T]](len(result))
	for _, item := range result {
		v := item.([]any)
		t := new(T)
		if v[0] != "" {
			catch.Try(sonic.Unmarshal([]byte(v[0].(string)), t))
		}
		vector.Append(ZElement[T]{
			Member: *t,
			Score:  v[1].(float64),
		})
	}
	return vector
}

func ZRangeRevObj[T any](key string, start, end int64) vec.Vec[T] {
	return zRangeObj[T](key, start, end, true)
}

func ZRangeObj[T any](key string, start, end int64) vec.Vec[T] {
	return zRangeObj[T](key, start, end, false)
}

func ZRangeRevByScoreObj[T any](key string, start, end int64) vec.Vec[ZElement[T]] {
	return zRangeByScoreObj[T](key, start, end, true)
}

func ZRangeByScoreObj[T any](key string, start, end int64) vec.Vec[ZElement[T]] {
	return zRangeByScoreObj[T](key, start, end, false)
}

func zRange(key string, start, end int64, rev bool) vec.Vec[string] {
	args := make([]any, 0, 6)
	if start < 0 {
		args = append(args, "zrange", key, "-inf", end)
	} else {
		args = append(args, "zrange", key, start, end)
	}
	if rev {
		args = append(args, "rev")
	}
	args = append(args, "byscore")
	result := catch.Try1(client.Do(context.Background(), args...).Slice())
	vector := vec.New[string](len(result))
	for _, item := range result {
		vector.Append(item.(string))
	}
	return vector
}

func zRangeByScore(key string, start, end int64, rev bool) vec.Vec[ZElement[string]] {
	args := make([]any, 0, 7)
	if start < 0 {
		args = append(args, "zrange", key, "-inf", end)
	} else {
		args = append(args, "zrange", key, start, end)
	}
	if rev {
		args = append(args, "rev")
	}

	args = append(args, "withscores", "byscore")
	result := catch.Try1(client.Do(context.Background(), args...).Slice())
	vector := vec.New[ZElement[string]](len(result))
	for _, item := range result {
		v := item.([]any)
		vector.Append(ZElement[string]{
			Member: v[0].(string),
			Score:  v[1].(float64),
		})
	}
	return vector
}

func ZRange(key string, start, end int64) vec.Vec[string] {
	return zRange(key, start, end, false)
}
func ZRangeRev(key string, start, end int64) vec.Vec[string] {
	return zRange(key, start, end, true)
}
func ZRangeByScore(key string, start, end int64) vec.Vec[ZElement[string]] {
	return zRangeByScore(key, start, end, false)
}

func ZRangeRevByScore(key string, start, end int64) vec.Vec[ZElement[string]] {
	return zRangeByScore(key, start, end, true)
}
