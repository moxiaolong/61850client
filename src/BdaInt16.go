package src

type BdaInt16 struct {
	BasicDataAttribute
}

func NewBdaInt16(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt16 {
	return &BdaInt16{}
}
