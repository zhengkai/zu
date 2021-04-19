package zu

import (
	"strings"
	"testing"
)

func TestFetchURL(t *testing.T) {

	UserAgent = `TestZU`

	ab, _ := FetchURL(`http://ifconfig.io/`)

	if !strings.Contains(string(ab), `TestZU`) {
		t.Error(`custom user agent failed`)
	}
}
