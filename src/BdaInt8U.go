package src

type BdaInt8U struct {
	BasicDataAttribute
	value int
}

func (i *BdaInt8U) setDefault() {
	i.value = 0
}
func NewBdaInt8U(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt8U {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT8U

	b := &BdaInt8U{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
