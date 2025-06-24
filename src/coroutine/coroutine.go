package coroutine

import "github.com/ihangsen/common/src/log"

// Launch 禁止在程序中直接使用go关键字，若需要用这个代替
func Launch(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Zap.Error(err)
			}
		}()
		fn()
	}()
}
