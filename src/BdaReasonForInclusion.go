package src

type BdaReasonForInclusion struct {
	BdaBitString
}

func NewBdaReasonForInclusion(objectReference *ObjectReference) *BdaReasonForInclusion {
	bitString := NewBdaBitString(objectReference, "", "", 7, false, false)
	b := &BdaReasonForInclusion{}
	b.BdaBitString = *bitString
	b.basicType = REASON_FOR_INCLUSION
	b.setDefault()
	return b
}
