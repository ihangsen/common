package res

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/log"
	"github.com/ihangsen/common/src/status"
	"github.com/ihangsen/common/src/utils/encipher"
	"github.com/ihangsen/common/src/utils/rawJson"
)

var (
	okJson = rawJson.RawJson(fmt.Sprintf(`{"code":%d}`, status.OK))
	ok     = func() []byte {
		bytes := rawJson.RawJson(fmt.Sprintf(`{"code":%d}`, status.OK))
		encipher.Encrypt(bytes)
		return bytes
	}()
	env string
)

func Init(env0 string) {
	env = env0
}

func Ok00(ctx *fiber.Ctx) error {
	return ctx.JSON(okJson)
}

func Ok0(ctx *fiber.Ctx) error {
	log.Zap.Info("请求完成")
	switch env {
	case "dev":
		return Ok00(ctx)
	case "test", "pro":
		return ctx.Send(ok)
	default:
		panic("请输入环境变量")
	}
}

func Ok1[T any](ctx *fiber.Ctx, data T) error {
	switch env {
	case "dev":
		return Ok11[T](ctx, data)
	case "test", "pro":
		log.Zap.Infof("响应数据：\n%+v", string(catch.Try1(sonic.Marshal(data))))
		bytes := catch.Try1(sonic.Marshal(Ok[T]{Code: status.OK, Data: data}))
		//log.Zap.Infof("响应加密后：\n%+v", bytes)
		encipher.Encrypt(bytes)
		return ctx.Send(bytes)
	default:
		panic("请输入环境变量")
	}
}

func Ok11[T any](ctx *fiber.Ctx, data T) error {
	log.Zap.Infof("响应数据：\n%+v", string(catch.Try1(sonic.Marshal(data))))
	return ctx.JSON(Ok[T]{Code: status.OK, Data: data})
}

func Ok111[T any](ctx *fiber.Ctx, data T) error {
	log.Zap.Infof("响应数据：\n%+v", string(catch.Try1(sonic.Marshal(data))))
	return ctx.JSON(data)
}

type Ok[T any] struct {
	Code uint8 `json:"code"`
	Data T     `json:"data"`
}
