package redis

import (
	"fmt"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/log"
	"github.com/ihangsen/common/src/redis"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestListObj(t *testing.T) {
	log.Init("dev")
	redis.Init(&redis.Redis{
		Addr:         "121.196.231.160:6379",
		Username:     "",
		Password:     "tTB9FcdLCCYK8jnB",
		Db:           1,
		PoolSize:     10,
		MinIdleConns: 1,
		MaxIdleConns: 10,
	})
	key := "list_test_obj"
	redis.LPushObj(key, vec.Vec[User]{User{
		Name: "1",
		Age:  1,
	}, User{
		Name: "2",
		Age:  2,
	}})
	obj := redis.LRangeObj[User](key, 0, -1)
	fmt.Println(obj)
}

func TestList(t *testing.T) {
	log.Init("dev")
	redis.Init(&redis.Redis{
		Addr:         "121.196.231.160:6379",
		Username:     "",
		Password:     "tTB9FcdLCCYK8jnB",
		Db:           1,
		PoolSize:     10,
		MinIdleConns: 1,
		MaxIdleConns: 10,
	})
	key := "list_test"
	redis.LPush[int](key, vec.Vec[int]{1, 2})
	obj := redis.LRange[string](key, 0, -1)
	fmt.Println(obj)
}

func TestLLen(t *testing.T) {
	log.Init("dev")
	redis.Init(&redis.Redis{
		Addr:         "121.196.231.160:6379",
		Username:     "",
		Password:     "tTB9FcdLCCYK8jnB",
		Db:           1,
		PoolSize:     10,
		MinIdleConns: 1,
		MaxIdleConns: 10,
	})
	key := "list_test"
	obj := redis.LLen(key)
	fmt.Println(obj)
}

func TestLRem(t *testing.T) {
	log.Init("dev")
	redis.Init(&redis.Redis{
		Addr:         "121.196.231.160:6379",
		Username:     "",
		Password:     "tTB9FcdLCCYK8jnB",
		Db:           1,
		PoolSize:     10,
		MinIdleConns: 1,
		MaxIdleConns: 10,
	})
	key := "list_test"
	redis.LRem[int](key, 1)
}

func TestLRemObj(t *testing.T) {
	log.Init("dev")
	redis.Init(&redis.Redis{
		Addr:         "121.196.231.160:6379",
		Username:     "",
		Password:     "tTB9FcdLCCYK8jnB",
		Db:           1,
		PoolSize:     10,
		MinIdleConns: 1,
		MaxIdleConns: 10,
	})
	key := "list_test_obj"
	redis.LRemObj[User](key, User{
		Name: "1",
		Age:  1,
	})
}
