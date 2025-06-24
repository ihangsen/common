package trans

import (
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
)

func Obj2Map[T any](obj any) map[string]T {
	data := catch.Try1(sonic.Marshal(obj))
	temp := new(map[string]T)
	catch.Try(sonic.Unmarshal(data, temp))
	return *temp
}

func Map2Obj[T any](m map[string]any) T {
	data := catch.Try1(sonic.Marshal(m))
	t := new(T)
	catch.Try(sonic.Unmarshal(data, t))
	return *t
}
