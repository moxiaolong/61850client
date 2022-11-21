package src

type BdaInt64 struct {
	BasicDataAttribute
}

func NewBdaInt64(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt64 {
	return &BdaInt64{}
}
