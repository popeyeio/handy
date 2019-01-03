package handy

import (
	"time"
)

const (
	DateGMT       = "Mon, 02 Jan 2006 15:04:05 GMT"
	DateLunar     = "2006-01-02 15:04:05"
	DateTZ        = "2006-01-02T15:04:05Z"
	DateLunarNano = "2006-01-02 15:04:05.000000000"
)

var (
	now = time.Now
)

func Second() int64 {
	return now().Unix()
}

func Millisecond() int64 {
	return now().UnixNano() / 1e6
}

func Microsecond() int64 {
	return now().UnixNano() / 1e3
}

func Nanosecond() int64 {
	return now().UnixNano()
}
