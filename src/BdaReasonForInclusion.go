package src

type BdaReasonForInclusion struct {
	BdaBitString
}

func NewBdaReasonForInclusion(*ObjectReference) *BdaReasonForInclusion {
	return &BdaReasonForInclusion{}
}
