package src

type BdaInt64 struct {
	BasicDataAttribute
	value int
}

func (i *BdaInt64) setDefault() {
	i.value = 0
}

func NewBdaInt64(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt64 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT64

	b := &BdaInt64{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
