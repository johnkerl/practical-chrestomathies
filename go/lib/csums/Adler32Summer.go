package csums

import (
	"fmt"
	"hash"
	"hash/adler32"
)

type Adler32Summer struct {
	// extends Checksummer is not needed in go.  just implement the functions.
	hash32 hash.Hash32
}

// scope is at package level.
// it would make sense to call it New if there were only one type in
// our package. but there are more than that.
// xxx to do:  how to hook into make()?
func NewAdler32Summer() *Adler32Summer {
	return &Adler32Summer{adler32.New()}
}

func (summer *Adler32Summer) Start() {
	summer.hash32.Reset()
}

func (summer *Adler32Summer) Accumulate(bytes []byte, n int) {
	// xxx what if n<len ... fix me ...
	if n < len(bytes) {
		// xxx n int, err error
		summer.hash32.Write(bytes[0:n])
	} else {
		summer.hash32.Write(bytes)
	}
}

func (summer *Adler32Summer) Finish() {
}

func (summer Adler32Summer) GetStringState() (state string) {
	return fmt.Sprintf("0x%08x", summer.hash32.Sum32())
}

func (summer Adler32Summer) GetStringSum() (sum string) {
	return fmt.Sprintf("0x%08x", summer.hash32.Sum32())
}
