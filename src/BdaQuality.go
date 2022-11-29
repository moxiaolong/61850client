package src

type BdaQuality struct {
	BdaBitString
	value  []byte
	mirror *BdaQuality
}

func (b *BdaQuality) copy() ModelNodeI {
	quality := NewBdaQuality(b.ObjectReference, b.Fc, b.sAddr, b.qchg)
	valueCopy := make([]byte, 0)
	copy(valueCopy, b.value)

	quality.value = valueCopy
	if b.mirror == nil {
		quality.mirror = b
	} else {
		quality.mirror = b.mirror
	}
	return quality
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

func (b *BdaQuality) GetValueString() string {
	if (b.value[0] & 0xC0) == 0xC0 {
		return "QUESTIONABLE"
	}

	if (b.value[0] & 0x80) == 0x80 {
		return "RESERVED"
	}

	if (b.value[0] & 0x40) == 0x40 {
		return "INVALID"
	}

	return "GOOD"

}
