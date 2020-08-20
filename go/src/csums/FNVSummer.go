package csums

import (
	"fmt"
	"hash"
	"hash/fnv"
)

type FNVSummer struct {
	// extends Checksummer is not needed in go.  just impl the fncs.
	hash32 hash.Hash32
	hash64 hash.Hash64
}

// scope is at package level.
// it would make sense to call it New if there were only one type in
// our package. but there are more than that.
// xxx to do:  how to hook into make()?
func NewFNVSummer(name string) *FNVSummer {
	// xxx nil, err?
	if name == "fnv32" {
		return &FNVSummer{fnv.New32(), nil}
	}
	if name == "fnv32a" {
		return &FNVSummer{fnv.New32a(), nil}
	}
	if name == "fnv64" {
		return &FNVSummer{nil, fnv.New64()}
	}
	if name == "fnv64a" {
		return &FNVSummer{nil, fnv.New64a()}
	}
	return nil
}

func (summer *FNVSummer) Start() {
	if summer.hash32 != nil {
		summer.hash32.Reset()
	}
	if summer.hash64 != nil {
		summer.hash64.Reset()
	}
}

func (summer *FNVSummer) Accumulate(bytes []byte, n int) {

	// xxx what if n<len ... fix me ...
	if summer.hash32 != nil {
		if n < len(bytes) {
			// xxx n int, err error
			summer.hash32.Write(bytes[0:n])
		} else {
			summer.hash32.Write(bytes)
		}
	}

	// xxx what if n<len ... fix me ...
	if summer.hash64 != nil {
		if n < len(bytes) {
			// xxx n int, err error
			summer.hash64.Write(bytes[0:n])
		} else {
			summer.hash64.Write(bytes)
		}
	}

}

func (summer *FNVSummer) Finish() {
}

func (summer FNVSummer) GetStringState() (state string) {
	return summer.GetStringSum()
}

func (summer FNVSummer) GetStringSum() (sum string) {
	var s string
	if summer.hash32 != nil {
		s = fmt.Sprintf("0x%08x", summer.hash32.Sum32())
	} else {
		s = fmt.Sprintf("0x%016x", summer.hash64.Sum64())
	}
	return s
}
