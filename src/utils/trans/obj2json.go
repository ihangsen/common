package trans

import (
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
)

func Obj2Json(obj any) string {
	data := catch.Try1(sonic.Marshal(obj))
	return string(data)
}

func Json2Obj[T any](m string) T {
	t := new(T)
	catch.Try(sonic.Unmarshal([]byte(m), t))
	return *t
}
