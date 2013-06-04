package csums

type Checksummer interface {
	Start()
	Accumulate(bytes []byte, len int)
	Finish()
	GetStringState() string
	GetStringSum()   string
}
