package handy

import (
	"reflect"
	"runtime"

	"github.com/valyala/fastrand"
)

func GetFuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func Choose(weights []uint32) (index int, ok bool) {
	if len(weights) == 0 {
		return
	}

	total := uint32(0)
	for _, weight := range weights {
		total += weight
	}

	n := fastrand.Uint32n(total) + 1
	for i, w := range weights {
		n -= w
		if n <= 0 {
			return i, true
		}
	}
	return
}
