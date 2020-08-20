package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s {one or more times in milliseconds since the epoch.",
			os.Args[0])
		os.Exit(1)
	}
	for _, arg := range args {
		// strconv.Atoi() doesn't do longs.
		var ms int64
		_, err := fmt.Sscanf(arg, "%d", &ms)
		if err == nil {
			s := ms / 1000
			ns := (ms % 1000) * 1000 * 1000
			t := time.Unix(s, ns)
			fmt.Println(t)
		} else {
			fmt.Fprintf(os.Stderr, "%s: can't parse \"%s\" as millis.\n",
				os.Args[0], arg)
		}
	}
}
