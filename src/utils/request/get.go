package request

import (
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/dict"
	"io"
	"net/http"
	"strings"
)

func Get0[T any](url string) *T {
	body := catch.Try1(http.Get(url)).Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	t := new(T)
	catch.Try(sonic.Unmarshal(catch.Try1(io.ReadAll(body)), t))
	return t
}

func Get1[T any](url0 string, params dict.Dict[string, string]) *T {
	builder := strings.Builder{}
	builder.WriteString(url0)
	if params.Len() > 0 {
		builder.WriteByte('?')
		flag := true
		params.ForEach(func(fieldName string, value string) {
			if flag {
				flag = false
			} else {
				builder.WriteByte('&')
			}
			builder.WriteString(fieldName)
			builder.WriteByte('=')
			builder.WriteString(value)
		})
	}
	body := catch.Try1(http.Get(builder.String())).Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	t := new(T)
	catch.Try(sonic.Unmarshal(catch.Try1(io.ReadAll(body)), t))
	return t
}
