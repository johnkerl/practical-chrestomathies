// go build cat.go
// ./cat myfile.txt
// -- or --
// go run cat.go myfile.txt

package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"io"
)

// ----------------------------------------------------------------
func main() {
	args := os.Args[1:]

	istream, err := argf(args)
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

// ----------------------------------------------------------------
func argf(filenames []string) (io.Reader, error) {
	if len(filenames) == 0 {
		return os.Stdin, nil
	} else {
		readers := make([]io.Reader, len(filenames))
		for i, filename := range(filenames) {
			handle, err := os.Open(filename)
			if err == nil {
				readers[i] = handle
			} else {
				return nil, err
			}
		}
		return io.MultiReader(readers...), nil
	}
}
