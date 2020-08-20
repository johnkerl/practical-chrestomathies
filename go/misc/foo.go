// go build hello.go
// ./hello a b c
package main

import (
	"fmt"
	"time"
)

func main() {
	ms := int64(1368847773000)
	s := ms / 1000
	ns := (ms % 1000) * 1000 * 1000
	t := time.Unix(s, ns)
	fmt.Println(t)
}
