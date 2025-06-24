package delay_task

import (
	"github.com/ihangsen/common/src/collection/dict"
	"time"
)

var (
	tasks = new(dict.SyncDict[uint32, *time.Timer])
	idInc = uint32(0)
)

func Add(delayTime time.Duration, fn func()) uint32 {
	idInc++
	fn0 := func() {
		defer tasks.LoadAndDelete(idInc)
		fn()
	}
	tasks.Store(idInc, time.AfterFunc(delayTime, fn0))
	return idInc
}

func Stop(id uint32) {
	tasks.LoadAndDelete(id).Map(func(t *time.Timer) { t.Stop() })
}

func Reset(id uint32, delayTime time.Duration) {
	tasks.Load(id).Map(func(t *time.Timer) { t.Reset(delayTime) })
}
