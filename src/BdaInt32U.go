package src

type BdaInt32U struct {
	BasicDataAttribute
	value int
}

func (i *BdaInt32U) setDefault() {
	i.value = 0
}
func NewBdaInt32U(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt32U {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT32U

	b := &BdaInt32U{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
