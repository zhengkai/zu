package main

import (
	"fmt"
	"strings"
	"zu"
)

var j = fmt.Println

func main() {

	j(zu.TS())

	j(zu.DumpHex([]byte{1, 2, 3, 4, 5}, 3))
	j(zu.DumpHex([]byte{1, 2, 3, 4, 5}, 10))

	ab, err := zu.FetchURL(`http://ifconfig.io/ip`)
	j(`fetch url:`, strings.TrimSpace(string(ab)), err)

	file := `/etc/passwd`
	j(`file`, zu.FileExists(file), file)

	file = `/etc/passwd-1283mcnakwk`
	j(`file`, zu.FileExists(file), file)

	json()
}

func json() {

	a := struct {
		Name string `json:"name"`
	}{
		Name: "hello",
	}

	j(zu.JSON(a))
}
