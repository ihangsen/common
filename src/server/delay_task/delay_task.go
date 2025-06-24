package delay_task

import (
	"github.com/ihangsen/common/src/collection/dict"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/coroutine"
	"github.com/ihangsen/common/src/server/redis"
	"sync"
	"time"
)

const (
	key      = "delayTask"
	Order    = 1
	LootRoom = 2
)

var (
	once     sync.Once
	callback = dict.Dict[uint8, func(task Task)]{}
)

type Task struct {
	OutTradeNo string `json:"outTradeNo,omitempty"`
	ChatRoomId uint64 `json:"chatRoomId,omitempty"`
	Type       uint8  `json:"type,omitempty"`
}

func Add(expire int64, task Task) {
	redis.ZAddObj[Task](key, vec.Of(redis.ZElement[Task]{
		Score:  float64(expire),
		Member: task,
	}))
}

// Callback 调用之前，必须先配置执行回调
func Callback(type0 uint8, fn func(task Task)) {
	callback.Store(type0, fn)
}

func Remove(task Task) {
	redis.ZRemObj[Task](key, vec.Of(task))
}

func Init() {
	once.Do(func() {
		coroutine.Launch(func() {
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					elements := redis.ZRangeByScoreObj[Task](key, -1, time.Now().UnixMilli())
					if elements.Len() <= 0 {
						continue
					}
					for _, element := range elements {
						task := element.Member
						callback.Load(task.Type).MapOrElse(func() {
							Remove(task)
						}, func(fn func(task Task)) {
							fn(task)
						})
					}
				}
			}
		})
	})
}
