package redis

import (
	"context"
	"github.com/ihangsen/common/src/catch"
	"github.com/redis/go-redis/v9"
)

var TestScript = redis.NewScript(`
local dayRichLen = tonumber(ARGV[1])
for i = 1, dayRichLen do
    local index = i * 2
	local score = ARGV[index]
    local member = ARGV[index + 1]
    redis.call("zincrby", KEYS[i], score, member)
end
return 1
`)

func Lua(script *redis.Script, keys []string, values []any) {
	catch.Try(script.Run(context.Background(), client, keys, values).Err())
}

func Lua1[T any](script *redis.Script, keys []string, values []any) T {
	return catch.Try1(script.Run(context.Background(), client, keys, values).Result()).(T)
}
