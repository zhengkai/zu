package zu

import (
	"io/ioutil"
	"net/http"
)

// UserAgent for FetchURL
var UserAgent = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.57`

// FetchURL ...
func FetchURL(url string) (ab []byte, err error) {

	req, err := http.NewRequest(`GET`, url, nil)
	if err != nil {
		return
	}
	req.Header.Set(`User-Agent`, UserAgent)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	ab, err = ioutil.ReadAll(res.Body)
	res.Body.Close()

	return
}
