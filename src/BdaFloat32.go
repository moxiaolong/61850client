package src

type BdaFloat32 struct {
	BasicDataAttribute
	value []byte
}

func (i *BdaFloat32) setDefault() {
	i.value = []byte{8, 0, 0, 0, 0}
}
func NewBdaFloat32(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaFloat32 {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = FLOAT32

	b := &BdaFloat32{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
