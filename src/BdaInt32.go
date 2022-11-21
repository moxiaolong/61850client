package src

type BdaInt32 struct {
	BasicDataAttribute
}

func NewBdaInt32(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt32 {
	return &BdaInt32{}
}
