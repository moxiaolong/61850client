package src

type BdaInt128 struct {
	BasicDataAttribute
}

func NewBdaInt128(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt128 {
	return &BdaInt128{}
}
