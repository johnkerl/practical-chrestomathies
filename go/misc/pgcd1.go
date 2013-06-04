package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gcd(a int, b int) (g int) {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	done := false
	for !done {
		r := a % b
		if r == 0 {
			done = true
		} else {
			a = b
			b = r
		}
	}

	return b
}

// -----------------------------------------------------------------------------
func main() {
	var nmin int = 200
	var nmax int = 10000
	var nsamp int = 100000
	var nreps int = 5

	if len(os.Args) == 2 {
		nsamp, _ = strconv.Atoi(os.Args[1])
	} else if len(os.Args) == 3 {
		nmax, _ = strconv.Atoi(os.Args[1])
		nsamp, _ = strconv.Atoi(os.Args[2])
	} else if len(os.Args) == 4 {
		nmin, _ = strconv.Atoi(os.Args[1])
		nmax, _ = strconv.Atoi(os.Args[2])
		nsamp, _ = strconv.Atoi(os.Args[3])
	}

	rand.Seed(time.Now().UnixNano() ^ int64(os.Getpid()))
	for n := nmin; n <= nmax; n++ {
		for j := 1; j <= nreps; j++ {
			count := 0
			for i := 1; i <= nsamp; i++ {
				u := 1 + int(rand.Int31n(int32(n)))
				v := 1 + int(rand.Int31n(int32(n)))
				g := gcd(u, v)
				if g == 1 {
					count += 1
				}
				//fmt.Println("--",u, v, g)
			}
			//fmt.Println(n, count, nsamp)
			// p := (1.0*count)/nsamp <---- == 0
			p := float64(count) / float64(nsamp)
			fmt.Printf("%6d %10.6f\n", n, p)
		}
	}
}
