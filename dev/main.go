package main

import (
	"fmt"
	"strings"
	"time"
	"zu"
	"zu/speedlimit"
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

	speedLimit()
}

func json() {

	a := struct {
		Name string `json:"name"`
	}{
		Name: "hello",
	}

	j(zu.JSON(a))
}

func speedLimit() {
	sl, _ := speedlimit.New(time.Second*2, 10)
	for i := 0; i < 12; i++ {
		j(`speedlimit`, i, sl.CanAdd())
	}
	time.Sleep(time.Second * 3)
	for i := 0; i < 12; i++ {
		j(`speedlimit`, i, sl.CanAdd())
	}
}
