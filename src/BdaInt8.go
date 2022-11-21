package src

type BdaInt8 struct {
	BasicDataAttribute
	value byte
}

func (i *BdaInt8) setDefault() {
	i.value = 0
}

func NewBdaInt8(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt8 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT8

	b := &BdaInt8{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
