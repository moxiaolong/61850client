package src

type BdaInt16 struct {
	BasicDataAttribute
	value int
}

func (i *BdaInt16) setDefault() {
	i.value = 0
}

func NewBdaInt16(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt16 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT16

	b := &BdaInt16{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
