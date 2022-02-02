package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, r := range []rune{'❦', '𐡷', '𐡸', '𐫱'} {
		fmt.Printf("%c %s\n", r, strconv.QuoteRuneToASCII(r))
	}

	// ❦ '\u2766'
	// 𐡷 '\U00010877'
	// 𐡸 '\U00010878'
	// 𐫱 '\U00010af1'

	// Note the lowercase \u for 4-digit sequences and the uppercase \U for
	// 8-digit sequences.
	fmt.Printf("\u2766\n")
	fmt.Printf("\U00010877\n")
	fmt.Printf("\U00010878\n")
	fmt.Printf("\U00010af1\n")

}
