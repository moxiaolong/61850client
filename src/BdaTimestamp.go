package src

type BdaTimestamp struct {
	BasicDataAttribute
	value []byte
}

func (t *BdaTimestamp) setDefault() {
	t.value = make([]byte, 8)
}

func NewBdaTimestamp(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaTimestamp {

	b := &BdaTimestamp{BasicDataAttribute: *NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)}
	b.basicType = TIMESTAMP
	b.setDefault()
	return b
}
