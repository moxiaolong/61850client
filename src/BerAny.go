package src

type BerAny struct {
	value []byte
}

func NewBerAny(value []byte) *BerAny {
	return &BerAny{value: value}
}
