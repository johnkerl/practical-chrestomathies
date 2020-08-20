package csums

import (
	"fmt"
)

type SimpleSummer struct {
	// extends Checksummer is not needed in go.  just impl the fncs.
	byteCount uint64
	byteSum   uint64
}

func (simpleSummer *SimpleSummer) Start() {
	simpleSummer.byteCount = 0
	simpleSummer.byteSum = 0
}

func (simpleSummer *SimpleSummer) Accumulate(bytes []byte, n int) {
	simpleSummer.byteCount += uint64(n)
	for i := 0; i < n; i++ {
		simpleSummer.byteSum += uint64(bytes[i])
	}
}

func (simpleSummer *SimpleSummer) Finish() {
}

func (simpleSummer SimpleSummer) GetStringState() (state string) {
	return fmt.Sprintf("%016x_%016x", simpleSummer.byteCount, simpleSummer.byteSum)
}

func (simpleSummer SimpleSummer) GetStringSum() (sum string) {
	return fmt.Sprintf("%016x_%016x", simpleSummer.byteCount, simpleSummer.byteSum)
}
