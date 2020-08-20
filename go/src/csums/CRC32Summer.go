package csums

import (
	"fmt"
	"hash"
	"hash/crc32"
)

type CRC32Summer struct {
	// extends Checksummer is not needed in go.  just impl the fncs.
	hash32 hash.Hash32
}

// scope is at package level.
// it would make sense to call it New if there were only one type in
// our package. but there are more than that.
// xxx to do:  how to hook into make()?
func NewCRC32Summer() *CRC32Summer {
	return &CRC32Summer{crc32.NewIEEE()}
}

func (summer *CRC32Summer) Start() {
	summer.hash32.Reset()
}

func (summer *CRC32Summer) Accumulate(bytes []byte, n int) {
	// xxx what if n<len ... fix me ...
	if n < len(bytes) {
		// xxx n int, err error
		summer.hash32.Write(bytes[0:n])
	} else {
		summer.hash32.Write(bytes)
	}
}

func (summer *CRC32Summer) Finish() {
}

func (summer CRC32Summer) GetStringState() (state string) {
	return fmt.Sprintf("0x%08x", summer.hash32.Sum32())
}

func (summer CRC32Summer) GetStringSum() (sum string) {
	return fmt.Sprintf("0x%08x", summer.hash32.Sum32())
}
