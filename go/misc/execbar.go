package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//pcmd := exec.Command("bash", "-c", "echo A; sleep 2; echo B")
	pcmd := exec.Command("bash", "-c", strings.Join(os.Args[1:], "\n"))

	o, oerr := pcmd.StdoutPipe()
	if oerr != nil {
		log.Fatal(oerr)
	}

	serr := pcmd.Start()
	if serr != nil {
		log.Fatal(serr)
	}

	reader := bufio.NewReader(o)
	eof := false
	for !eof {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Print(line)
		}
	}
}
