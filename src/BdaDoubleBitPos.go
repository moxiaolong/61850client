package src

type BdaDoubleBitPos struct {
	BdaBitString
}

func NewBdaDoubleBitPos(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaDoubleBitPos {
	super := NewBdaBitString(objectReference, fc, sAddr, 2, dchg, dupd)
	super.basicType = DOUBLE_BIT_POS
	b := &BdaDoubleBitPos{BdaBitString: *super}
	b.setDefault()
	return b
}
