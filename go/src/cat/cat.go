// From /path/to/practical-chrestomathies/go:
// export GOPATH=$(pwd)
// go build cat

package main

import (
	"argf"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// ----------------------------------------------------------------
func main() {
	args := os.Args[1:]

	istream, err := argf.Open(args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = cat(istream)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// ----------------------------------------------------------------
func cat(istream io.Reader) error {

	reader := bufio.NewReader(istream)
	eof := false

	for !eof {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			log.Println(err)
			return err
		} else {
			fmt.Print(line)
		}
	}

	return nil
}
