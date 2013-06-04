package main
// This program just runs until it overflows a 32-bit counter.  It takes a couple seconds
// on my laptop.
// The experiment here is to see how small the executable is when not importing anything.
func main() {
	var i uint32 = 1
	for i != 0 {
		i++
	}
}
