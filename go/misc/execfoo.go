package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//out, err := exec.Command("date").Output()
	//out, err := exec.Command("bash", "-c", "'date'").Output()
	//out, err := exec.Command("ls").Output()
	//out, err := exec.Command("bash", "-c", "ls|wc -l").Output()
	out, err := exec.Command("bash", "-c", "echo A; sleep 2; echo B").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output:\n")
	fmt.Printf("%s", out)
}
