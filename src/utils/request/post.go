package request

import (
	"bytes"
	"github.com/bytedance/sonic"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/log"
	"io"
	"net/http"
)

func PostJson[T, P any](url string, params P, fn func(header http.Header)) *T {
	buffer := bytes.NewBuffer(catch.Try1(sonic.Marshal(params)))
	request := catch.Try1(http.NewRequest(http.MethodPost, url, buffer))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	fn(request.Header)
	body := catch.Try1(http.DefaultClient.Do(request)).Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	t := new(T)
	bodyBytes := catch.Try1(io.ReadAll(body))
	log.Zap.Infof("body:%s", string(bodyBytes))
	catch.Try(sonic.Unmarshal(bodyBytes, t))
	return t
}

func PostFrom[T, P any](url string, params P) *T {
	buffer := bytes.NewBuffer(catch.Try1(sonic.Marshal(params)))
	request := catch.Try1(http.NewRequest(http.MethodPost, url, buffer))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body := catch.Try1(http.DefaultClient.Do(request)).Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	t := new(T)
	catch.Try(sonic.Unmarshal(catch.Try1(io.ReadAll(body)), t))
	return t
}
