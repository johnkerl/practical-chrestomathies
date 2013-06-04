// go build hello.go
// ./hello a b c
package main
import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	mcmd := strings.Join(os.Args[1:], " ")
	var pcmd *exec.Cmd = exec.Command("bash", "-c", mcmd)

	err := pcmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("done\n")

//	out, err := pcmd.Output()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Output:\n")
//	fmt.Printf("%s\n", out)

}
