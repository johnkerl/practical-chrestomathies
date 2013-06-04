// go build csum.go
// ./csum myfile.txt
// -- or --
// go run csum.go myfile.txt
// go run csum.go -- csum.go

// http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file

package main

import (
	// xxx not ";" sep since unused-import error ... we really need easy
	// import/not with "//".
	"csums"
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
	fmt.Fprintf(os.Stderr, "If no filenames are given, stdin is read.\n")
	flag.PrintDefaults()
	os.Exit(1)
}

// ----------------------------------------------------------------
func main() {
	// http://golang.org/pkg/flag
	pDoSpin   := flag.Bool("spin", false, "print running checksums to screen (default off)")
	// xxx string-set ... maybe do my own non-flag here?
	// xxx make a list-algos opt ...
	algoNames := strings.Join(csums.ChecksummerFactoryAlgoNames, ", ")
	pAlgo    := flag.String("algo", "crc64", "One of: " + algoNames + ".")

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	doSpin := *pDoSpin

	checksummer, err := csums.ChecksummerFactory(*pAlgo)
	if (checksummer == nil) {
		log.Fatal(err)
	}

	ok := true
	if len(args) == 0 {
		ok = csum("-", checksummer, doSpin) && ok
	} else {
		for _, arg := range args {
			// ok && count(arg): not called after error
			ok = csum(arg, checksummer, doSpin) && ok
		}
	}
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

// ----------------------------------------------------------------
func csum(fileName string, checksummer csums.Checksummer, doSpin bool) (ok bool) {

	sourceStream := os.Stdin
	if fileName != "-" {
		var err error
		if sourceStream, err = os.Open(fileName); err != nil {
			log.Println(err)
			return false
		}
	}

	bufferSize := 2048

	buffer := make([]byte, bufferSize)
	eof := false
	var nblocks uint64 = 0
	var nbytes  uint64 = 0

	checksummer.Start()

	if (doSpin) {
		fmt.Print("...")
		os.Stdout.Sync()
	}

	for !eof {
		numBytesRead, err := sourceStream.Read(buffer)
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			log.Println(err)
			if fileName != "-" {
				sourceStream.Close()
			}
			return false
		} else {
			nblocks ++
			nbytes += uint64(numBytesRead)

			checksummer.Accumulate(buffer, numBytesRead)

			if (doSpin && ((nblocks & 0x3ff) == 0)) {
				fmt.Print("\r"+checksummer.GetStringState())
				os.Stdout.Sync()
			}
		}
	}

	if fileName != "-" {
		sourceStream.Close()
	}

	checksummer.Finish()
	stringSum := checksummer.GetStringSum()

	if (doSpin) {
		fmt.Print("\r")
	}
	fmt.Printf("%s  %6d  %s\n", stringSum, nbytes, fileName)

	return true
}
