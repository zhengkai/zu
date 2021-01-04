package main

import (
	"fmt"
	"zu"
)

var j = fmt.Println

func main() {

	j(zu.TS())

	j(zu.DumpHex([]byte{1, 2, 3, 4, 5}, 3))
	j(zu.DumpHex([]byte{1, 2, 3, 4, 5}, 10))
}
