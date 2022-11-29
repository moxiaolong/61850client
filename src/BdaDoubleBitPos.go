package src

type BdaDoubleBitPos struct {
	BdaBitString
	mirror *BdaDoubleBitPos
}

func (s *BdaDoubleBitPos) copy() ModelNodeI {
	newCopy := NewBdaDoubleBitPos(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)
	valueCopy := make([]byte, 0)
	copy(valueCopy, s.value)
	newCopy.value = valueCopy
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func NewBdaDoubleBitPos(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaDoubleBitPos {
	super := NewBdaBitString(objectReference, fc, sAddr, 2, dchg, dupd)
	super.basicType = DOUBLE_BIT_POS
	b := &BdaDoubleBitPos{BdaBitString: *super}
	b.setDefault()
	return b
}

func (b *BdaDoubleBitPos) GetValueString() string {
	if (b.value[0] & 0xC0) == 0xC0 {
		return "BAD_STATE"
	}

	if (b.value[0] & 0x80) == 0x80 {
		return "ON"
	}

	if (b.value[0] & 0x40) == 0x40 {
		return "OFF"
	}

	return "INTERMEDIATE_STATE"
}
