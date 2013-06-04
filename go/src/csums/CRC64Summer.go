package csums

import(
	"fmt"
	"hash"
	"hash/crc64"
)

type CRC64Summer struct {
	// extends Checksummer is not needed in go.  just impl the fncs.
	hash64 hash.Hash64
}

// scope is at package level.
// it would make sense to call it New if there were only one type in
// our package. but there are more than that.
// xxx to do:  how to hook into make()?
func NewCRC64Summer() *CRC64Summer {
	var table *crc64.Table = crc64.MakeTable(crc64.ISO)
	return &CRC64Summer{crc64.New(table)}
}

func (summer *CRC64Summer) Start() {
	summer.hash64.Reset()
}

func (summer *CRC64Summer) Accumulate(bytes []byte, n int) {
	// xxx what if n<len ... fix me ...
	if (n < len(bytes)) {
		// xxx n int, err error
		summer.hash64.Write(bytes[0:n])
	} else {
		summer.hash64.Write(bytes)
	}
}

func (summer *CRC64Summer) Finish() {
}

func (summer CRC64Summer) GetStringState() (state string) {
	return fmt.Sprintf("0x%016x", summer.hash64.Sum64())
}

func (summer CRC64Summer) GetStringSum() (sum string) {
	return fmt.Sprintf("0x%016x", summer.hash64.Sum64())
}
