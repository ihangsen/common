package catch

import (
	"github.com/ihangsen/common/src/log"
)

func Try(err error) {
	if err != nil {
		log.Zap.Error(err)
		panic(err)
	}
}

func Try1[T0 any](t0 T0, e error) T0 {
	Try(e)
	return t0
}

func Try2[T0, T1 any](t0 T0, t1 T1, e error) (T0, T1) {
	Try(e)
	return t0, t1
}

func Try3[T0, T1, T2 any](t0 T0, t1 T1, t2 T2, e error) (T0, T1, T2) {
	Try(e)
	return t0, t1, t2
}

func Try4[T0, T1, T2, T3 any](t0 T0, t1 T1, t2 T2, t3 T3, e error) (T0, T1, T2, T3) {
	Try(e)
	return t0, t1, t2, t3
}

func Try5[T0, T1, T2, T3, T4 any](t0 T0, t1 T1, t2 T2, t3 T3, t4 T4, e error) (T0, T1, T2, T3, T4) {
	Try(e)
	return t0, t1, t2, t3, t4
}
