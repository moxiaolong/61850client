package src

type BerAny struct {
	value []byte
}

func (a *BerAny) encode(reverseOS *ReverseByteArrayOutputStream) int {
	reverseOS.write(a.value)
	return len(a.value)
}

func NewBerAny(value []byte) *BerAny {
	return &BerAny{value: value}
}
