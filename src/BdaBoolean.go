package src

type BdaBoolean struct {
	BasicDataAttribute
	value bool
}

func (i *BdaBoolean) setDefault() {
	i.value = false
}
func NewBdaBoolean(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaBoolean {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = BOOLEAN

	b := &BdaBoolean{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
