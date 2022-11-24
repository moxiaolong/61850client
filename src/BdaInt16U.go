package src

type BdaInt16U struct {
	BasicDataAttribute
	value int
}

func (i *BdaInt16U) setDefault() {
	i.value = 0
}
func NewBdaInt16U(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt16U {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT16U

	b := &BdaInt16U{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
