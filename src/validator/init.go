package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/i18n"
	"github.com/ihangsen/common/src/res"
	"regexp"
	"sync"
	"time"
)

var (
	validate = validator.New()
	once     sync.Once
)

const (
	regexpUsername = "^1([3-9])\\d{9}$"
	regexpPassword = "^[a-zA-Z\\d!@#$%^&*()_+={};':\"|,.<>?/]{7,16}$"
)

func Init() {
	once.Do(func() {
		validate.SetTagName("binding")
		catch.Try(validate.RegisterValidation("username", username))
		catch.Try(validate.RegisterValidation("password", password))
		catch.Try(validate.RegisterValidation("submeterTime", submeterTime))
	})
}

func Struct[T any](t *T) error {
	return validate.Struct(t)
}

var username validator.Func = func(fl validator.FieldLevel) bool {
	username, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if username == "" {
		return true
	}
	ok = catch.Try1(regexp.MatchString(regexpUsername, username))
	if !ok {
		res.Msg(i18n.Get.UserFormatErr)
	}
	return true
}

var password validator.Func = func(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if password == "" {
		return true
	}
	ok = catch.Try1(regexp.MatchString(regexpPassword, password))
	if !ok {
		res.Msg(i18n.Get.PwdFormatErr)
	}
	return ok
}

var submeterTime validator.Func = func(fl validator.FieldLevel) bool {
	year, ok := fl.Field().Interface().(int)
	if !ok {
		return false
	}
	now := time.Now()
	nowYear := now.Year()
	if year > nowYear || year < 2025 {
		res.Msg(i18n.Get.TimeErr)
	}
	return ok
}
