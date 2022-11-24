package src

type BdaInt32 struct {
	BasicDataAttribute
	value int
}

func (i *BdaInt32) setDefault() {
	i.value = 0
}

func NewBdaInt32(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt32 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT32

	b := &BdaInt32{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
