// go build bytehisto.go
// ./bytehisto myfile.txt
// -- or --
// go run bytehisto.go myfile.txt
// go run bytehisto.go -- bytehisto.go

// xxx cmt me esp. about the -c option
// xxx refactor me ...

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
	pDoCleanOnly := flag.Bool("c", false, "Only test for ASCII-printables only")
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	doCleanOnly := *pDoCleanOnly

	ok := true
	clean := true
	if len(args) == 0 {
		currOk, currClean := bytehistoDump("-", doCleanOnly)
		// &&= DNE?!?
		ok = ok && currOk
		clean = clean && currClean
	} else {
		for _, arg := range args {
			currOk, currClean := bytehistoDump(arg, doCleanOnly) // ok && count(arg): not called after error
			ok = ok && currOk
			clean = clean && currClean
		}
	}
	if doCleanOnly {
		ok = ok && clean
	}
	if ok {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

// ----------------------------------------------------------------
func bytehistoDump(sourceName string, doCleanOnly bool) (ok bool, clean bool) {

	sourceStream := os.Stdin
	if sourceName != "-" {
		var err error
		if sourceStream, err = os.Open(sourceName); err != nil {
			log.Println(err)
			return false, false
		}
	}

	bufferSize := 2048

	buffer := make([]byte, bufferSize)
	counts := make([]int64, 256)
	eof := false

	for !eof {
		numBytesRead, err := sourceStream.Read(buffer)
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			log.Println(err)
			if sourceName != "-" {
				sourceStream.Close()
			}
			return false, false
		} else {
			for i := 0; i < numBytesRead; i++ {
				counts[buffer[i]]++
			}
		}
	}

	// xxx defer func
	if sourceName != "-" {
		sourceStream.Close()
	}

	clean = true
	minCount := int64(1)
	maxCount := int64(0)
	minByte := int32(257)
	maxByte := int32(0)

	for i := 0; i < 256; i++ {
		if counts[i] < minCount {
			minCount = counts[i]
		}
		if counts[i] > maxCount {
			maxCount = counts[i]
		}
		if counts[i] != 0 {
			if minByte == 257 {
				minByte = int32(i)
			}
			maxByte = int32(i)
		}
	}

	var printName = sourceName
	if sourceName == "-" {
		printName = "[stdin]"
	}
	if !doCleanOnly {
		fmt.Printf("----------------------------------------------------------------\n")
		fmt.Printf("%s:\n", printName)
	}

	// xxx separate methods for file read, stats collect, table print ...

	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			n := i + 32*j
			// Go will print various Unicode things but for my purposes I want
			// to know if a text file is ASCII clean.  So I'm not using
			// strconv.IsPrint here.
			isprint := (n == 0x09) || (n == 0x0a) || ((n >= 0x20) && (n <= 0x7e))
			if counts[n] != 0 {
				clean = clean && isprint
			}

			if !doCleanOnly {
				if isprint && (n != ' ') && (n != '\t') {
					fmt.Printf("%2c", n)
				} else {
					fmt.Printf("%02x", n)
				}
				fmt.Printf(": %-5d", counts[n])
				if j == 7 {
					fmt.Printf(" ")
				}
			}
		}
		if !doCleanOnly {
			fmt.Printf("\n")
		}
	}
	if !doCleanOnly {
		if maxCount > 0 {
			fmt.Printf("Min byte : %-3d (0x%02x)  Max byte : %-3d (0x%02x)\n", minByte, minByte, maxByte, maxByte)
		}
		fmt.Printf("Min count: %-6d  Max count: %-6d\n", minCount, maxCount)
		fmt.Printf("\n")
	} else {
		fmt.Printf("%-5v %s\n", clean, printName)
	}

	return true, clean
}
