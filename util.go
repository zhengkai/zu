package zu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

// FetchURL ...
func FetchURL(url string) (ab []byte, err error) {

	res, err := http.Get(url)
	if err != nil {
		return
	}
	ab, err = ioutil.ReadAll(res.Body)
	res.Body.Close()

	return
}

// FileExists ...
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// JSON ...
func JSON(a interface{}) (s string) {
	ab, err := json.Marshal(a)
	if err != nil {
		s = fmt.Sprintf(`(json marshal err: %s)`, err)
		return
	}
	s = string(ab)
	return
}

// JSONPretty ...
func JSONPretty(a interface{}) (s string) {
	ab, err := json.MarshalIndent(a, ``, "\t")
	if err != nil {
		s = fmt.Sprintf(`(json marshal err: %s)`, err)
		return
	}
	s = string(ab)
	return
}
