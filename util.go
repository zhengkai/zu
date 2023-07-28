package zu

import (
	"encoding/json"
	"fmt"
)

// DumpHex dump first n bytes
func DumpHex(ab []byte, n int) string {
	out := ab
	if len(ab) > n {
		out = ab[:n]
	}
	return fmt.Sprintf(`%x`, out)
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
