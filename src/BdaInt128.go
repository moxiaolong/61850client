package src

type BdaInt128 struct {
	BasicDataAttribute
	value int
}

func (i *BdaInt128) setDefault() {
	i.value = 0
}

func NewBdaInt128(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt128 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT128

	b := &BdaInt128{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
