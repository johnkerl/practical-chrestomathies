// go build wc.go
// ./wc myfile.txt
// -- or --
// go run wc.go myfile.txt
// go run wc.go -- wc.go

package main

import (
	// xxx not ";" sep since unused-import error ... we really need easy
	// import/not with "//".
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ----------------------------------------------------------------
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] {filenames ...}\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "If no file name is given, or if filename is \"-\", stdin is used.\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

// ----------------------------------------------------------------
func main() {
	pCountLines := flag.Bool("l", false, "Count lines")
	pCountWords := flag.Bool("w", false, "Count words")
	pCountChars := flag.Bool("c", false, "Count chars")

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	countLines := *pCountLines
	countWords := *pCountWords
	countChars := *pCountChars
	if !countLines && !countWords && !countChars {
		countLines = true
		countWords = true
		countChars = true
	}

	ok := true
	if len(args) == 0 {
		ok = count("-", countLines, countWords, countChars) && ok
	} else {
		for _, arg := range args {
			ok = count(arg, countLines, countWords, countChars) && ok // ok && count(arg): not called after error
		}
	}
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

// ----------------------------------------------------------------
func count(sourceName string, countLines bool, countWords bool,
	countChars bool) (ok bool) {

	var numLines = 0
	var numWords = 0
	var numChars = 0

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
			// line = strings.TrimRight(line, "\n")

			// This is how to trim leading/trailing whitespace:
			//line = strings.TrimSpace(line)

			numLines++
			numChars += len(line)
			numWords += len(strings.Fields(line))
		}
	}

	if sourceName != "-" {
		sourceStream.Close()
	}

	if countLines {
		fmt.Printf(" %4d", numLines)
	}
	if countWords {
		fmt.Printf(" %4d", numWords)
	}
	if countChars {
		fmt.Printf(" %4d", numChars)
	}
	fmt.Println(" ", sourceName)

	return true
}
