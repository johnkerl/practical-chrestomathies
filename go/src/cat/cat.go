// go build cat.go
// ./cat myfile.txt
// -- or --
// go run cat.go myfile.txt
// go run cat.go -- cat.go

package main

import (
	// xxx not ";" sep since unused-import error ... we really need easy
	// import/not with "//".
	"fmt"
	"log"
	"os"
	"bufio"
	"io"
)

// ----------------------------------------------------------------
func main() {
	args := os.Args[1:]

	ok := true
	if len(args) == 0 {
		ok = cat("-") && ok
	} else {
		for _, arg := range args {
			ok = cat(arg) && ok // ok && cat(arg): not called after error
		}
	}
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

// ----------------------------------------------------------------
func cat(sourceName string) (ok bool) {
	sourceStream := os.Stdin
	if sourceName != "-" {
		var err error
		if sourceStream, err = os.Open(sourceName); err != nil {
			log.Println(err)
			return false
		}
	}

	reader := bufio.NewReader(sourceStream)
	eof := false

	for !eof {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			log.Println(err)
			if sourceName != "-" {
				sourceStream.Close()
			}
			return false
		} else {
			// This is how to do a chomp:
			//line = strings.TrimRight(line, "\n")
			fmt.Print(line)
		}
	}
	if sourceName != "-" {
		sourceStream.Close()
	}

	return true
}
