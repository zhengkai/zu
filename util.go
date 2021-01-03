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
