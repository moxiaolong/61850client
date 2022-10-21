package src

type BerAny struct {
	value []byte
}

func NewBerAny([]byte) *BerAny {
	return &BerAny{}
}
