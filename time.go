package zu

import "time"

// TS unixstamp in second
func TS() uint32 {
	return uint32(time.Now().Unix())
}

// MS unixstamp in millisecond
func MS() uint64 {
	return uint64(time.Now().UnixNano() / 1000000)
}

// TSFormat ...
func TSFormat[T int64 | uint64 | int | uint | int32 | uint32](ts T) string {
	if ts >= 2147483647 {
		ts = ts / 1000
	}
	return time.Unix(int64(ts), 0).Format(time.DateTime)
}
