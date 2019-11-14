package common

import "time"

func ToMillisecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
