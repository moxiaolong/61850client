package src

type BdaFloat64 struct {
	BasicDataAttribute
	value []byte
}

func NewBdaFloat64(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaFloat64 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = FLOAT64

	b := &BdaFloat64{BasicDataAttribute: *attribute}
	b.setDefault()
	return b

}

func (i *BdaFloat64) setDefault() {
	i.value = []byte{8, 0, 0, 0, 0}
}
