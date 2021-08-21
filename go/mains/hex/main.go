// From /path/to/practical-chrestomathies/go:
// export GOPATH=$(pwd)
// go build hex

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/johnkerl/practical-chrestomathies/go/lib/argf"
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
	pDoRaw := flag.Bool("r", false, "Count lines")

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	doRaw := *pDoRaw

	istream, err := argf.Open(args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = hexDump(istream, doRaw)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

// ----------------------------------------------------------------
func hexDump(sourceStream io.Reader, doRaw bool) error {

	bytesPerClump := 4
	clumpsPerLine := 4
	bufferSize := bytesPerClump * clumpsPerLine

	buffer := make([]byte, bufferSize)
	eof := false
	offset := 0

	for !eof {
		numBytesRead, err := io.ReadFull(sourceStream, buffer)
		if err == io.EOF {
			eof = true
			break
		}
		// io.ErrUnexpectedEOF is the normal case when the file size isn't an
		// exact multiple of our buffer size.
		if err != nil && err != io.ErrUnexpectedEOF {
			log.Println(err)
			return err
		}

		// Print offset "pre" part
		if !doRaw {
			fmt.Printf("%08x: ", offset)
		}

		// Print hex payload
		for i := 0; i < bufferSize; i++ {
			if i < numBytesRead {
				fmt.Printf("%02x ", buffer[i])
			} else {
				fmt.Printf("   ")
			}
			if (i % bytesPerClump) == (bytesPerClump - 1) {
				if (i > 0) && (i < bufferSize-1) {
					fmt.Printf(" ")
				}
			}
		}

		// Print ASCII-dump "post" part
		if !doRaw {
			fmt.Printf("|")

			for i := 0; i < numBytesRead; i++ {
				if buffer[i] >= 0x20 && buffer[i] <= 0x7e {
					fmt.Printf("%c", buffer[i])
				} else {
					fmt.Printf(".")
				}
			}
			for i := numBytesRead; i < bufferSize; i++ {
				fmt.Print(" ")
			}
			fmt.Printf("|")
		}

		// Print line end
		fmt.Printf("\n")

		offset += numBytesRead

	}

	return nil
}
