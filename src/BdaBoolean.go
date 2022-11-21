package src

type BdaBoolean struct {
	BasicDataAttribute
}

func NewBdaBoolean(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaBoolean {
	return &BdaBoolean{}
}
