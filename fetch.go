package zu

import (
	"io/ioutil"
	"net/http"
)

// UserAgent for FetchURL
var UserAgent = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36 Edg/89.0.774.77`

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
