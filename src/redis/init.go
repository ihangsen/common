package redis

import (
	"context"
	"errors"
	"github.com/ihangsen/common/src/log"
	"github.com/redis/go-redis/v9"
	"sync"
)

type Redis struct {
	Addr     string
	Username string
	Password string
	Db       int
	//DialTimeout     time.Duration // 建立新网络连接时的超时时间 默认5秒
	//ReadTimeout     time.Duration // 读超时 默认值，3秒
	//WriteTimeout    time.Duration // 写超时 默认值，3秒
	PoolSize     int // 连接池最大连接数量
	MinIdleConns int // 连接池保持的最小空闲连接数
	MaxIdleConns int // 连接池保持的最大空闲连接数，多余的空闲连接将被关闭
}

var client = new(redis.Client)
var once sync.Once

func Init(conf *Redis) {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:         conf.Addr,
			Username:     conf.Username,
			Password:     conf.Password,
			DB:           conf.Db,
			PoolSize:     conf.PoolSize,
			MinIdleConns: conf.MinIdleConns,
			MaxIdleConns: conf.MaxIdleConns,
		})
		pong := client.Ping(context.Background()).String()
		if pong != "ping: PONG" {
			panic("redis连接失败")
		} else {
			log.Zap.Info("redis连接成功")
		}
	})
}

func try(err error) bool {
	if err == nil {
		return true
	}
	if errors.Is(err, redis.Nil) {
		return false
	}
	panic(err)
}
