package trans

import (
	"github.com/ihangsen/common/src/catch"
	"time"
)

const (
	YyyyMmDdHhMmSs  = "2006-01-02 15:04:05"
	YyyyMmDdHhMmSsT = "2006-01-02T15:04:05+08:00"
	YyyyMmDd        = "2006-01-02"
)

func Str2Milli(format string, value string) int64 {
	return catch.Try1(time.ParseInLocation(format, value, time.Local)).UnixMilli()
}

func Milli2Str(format string, value int64) string {
	return time.UnixMilli(value).Format(format)
}
