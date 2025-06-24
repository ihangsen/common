package res

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ihangsen/common/src/i18n"
	"github.com/ihangsen/common/src/log"
	"github.com/ihangsen/common/src/status"
	"github.com/ihangsen/common/src/utils/encipher"
	"github.com/ihangsen/common/src/utils/rawJson"
	"net/http"
)

var (
	tokenValid       = Err{Code: status.TokenValid, Msg: i18n.Get.TokenValid}
	tooEarly         = Err{Code: status.TooEarly, Msg: i18n.Get.TooEarly}
	notLogin         = Err{Code: status.NotLogin, Msg: i18n.Get.NotLogin}
	refreshValid     = Err{Code: status.RefreshValid, Msg: i18n.Get.RefreshValid}
	PermissionDenied = Err{Code: status.PermissionDenied, Msg: i18n.Get.PermissionDenied}
	systemErrorJson  = rawJson.RawJson(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, status.SystemError, i18n.Get.SystemErr))
	systemError      = func() []byte {
		bytes := rawJson.RawJson(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, status.SystemError, i18n.Get.SystemErr))
		encipher.Encrypt(bytes)
		return bytes
	}()
)

func Msg(msg string) {
	panic(Err{Code: status.CustomCode, Msg: msg})
}

func Code(code uint8) {
	panic(Err{Code: code})
}

func Error(err Err) {
	panic(err)
}

func TokenValid() {
	panic(tokenValid)
}

func TooEarly() {
	panic(tooEarly)
}

func NotLogin() {
	panic(notLogin)
}

func RefreshValid() {
	panic(refreshValid)
}

type Err struct {
	Code uint8  `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func (r Err) Error() string {
	if r.Msg == "" {
		return fmt.Sprintf(`{"code":%d}`, r.Code)
	}
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, r.Code, r.Msg)
}

func ErrorHandler0(ctx *fiber.Ctx) error {
	defer func() {
		switch e := recover().(type) {
		case nil:
		case Err:
			log.Zap.Error(e)
			_ = ctx.JSON(e)
		default:
			log.Zap.Error(e)
			ctx.Status(http.StatusInternalServerError)
			_ = ctx.JSON(systemErrorJson)
		}
	}()
	_ = ctx.Next()
	return nil
}

func ErrorHandler(ctx *fiber.Ctx) error {
	log.Zap.Infof("Path: %s %s", ctx.Method(), ctx.Path())
	switch env {
	case "dev":
		return ErrorHandler0(ctx)
	case "test", "pro":
		defer func() {
			switch e := recover().(type) {
			case nil:
			case Err:
				log.Zap.Error(e)
				bytes := []byte(e.Error())
				encipher.Encrypt(bytes)
				_ = ctx.Send(bytes)
			default:
				log.Zap.Error(e)
				ctx.Status(http.StatusInternalServerError)
				_ = ctx.Send(systemError)
			}
		}()
		_ = ctx.Next()
		return nil
	default:
		panic("请输入环境变量")
	}
}
