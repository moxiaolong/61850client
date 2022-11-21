package src

type BdaQuality struct {
	BdaBitString
	value []byte
}

func NewBdaQuality(objectReference *ObjectReference, fc string, sAddr string, qchg bool) *BdaQuality {
	b := &BdaQuality{BdaBitString: *NewBdaBitString(objectReference, fc, sAddr, 13, qchg, false)}

	b.qchg = qchg
	b.basicType = QUALITY
	b.setDefault()
	return b
}

func (b *BdaQuality) setDefault() {
	b.value = []byte{0x00, 0x00}
}
