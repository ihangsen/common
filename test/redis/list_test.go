package redis

import (
	"fmt"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/redis"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestListObj(t *testing.T) {
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
