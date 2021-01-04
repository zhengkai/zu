package zu

import (
	"fmt"
	"time"
)

// TS unixstamp in second
func TS() uint32 {
	return uint32(time.Now().Unix())
}

// MS unixstamp in millisecond
func MS() uint64 {
	return uint64(time.Now().UnixNano() / 1000000)
}

// DumpHex dump first n bytes
func DumpHex(ab []byte, n int) string {
	out := ab
	if len(ab) > n {
		out = ab[:n]
	}
	return fmt.Sprintf(`%x`, out)
}
