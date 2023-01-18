package src

type BdaOptFlds struct {
	BdaBitString
	mirror *BdaOptFlds
}

func (f *BdaOptFlds) copy() ModelNodeI {
	newCopy := NewBdaOptFlds(f.ObjectReference, f.Fc)
	valueCopy := make([]byte, 0)
	copy(valueCopy, f.value)
	newCopy.value = valueCopy
	if f.mirror == nil {
		newCopy.mirror = f
	} else {
		newCopy.mirror = f.mirror
	}
	return newCopy
}

func (f *BdaOptFlds) isSequenceNumber() bool {
	return (f.value[0] & 0x40) == 0x40
}

func (f *BdaOptFlds) isReportTimestamp() bool {
	return (f.value[0] & 0x20) == 0x20
}

func (f *BdaOptFlds) isDataSetName() bool {
	return (f.value[0] & 0x08) == 0x08
}

func (f *BdaOptFlds) isBufferOverflow() bool {
	return (f.value[0] & 0x02) == 0x02
}

func (f *BdaOptFlds) isEntryId() bool {
	return (f.value[0] & 0x01) == 0x01
}

func (f *BdaOptFlds) isConfigRevision() bool {
	return (f.value[1] & 0x80) == 0x80
}

func (f *BdaOptFlds) isSegmentation() bool {
	return (f.value[1] & 0x40) == 0x40
}

func (f *BdaOptFlds) isDataReference() bool {
	return (f.value[0] & 0x04) == 0x04
}

func (f *BdaOptFlds) isReasonForInclusion() bool {
	return (f.value[0] & 0x10) == 0x10
}

func (f *BdaOptFlds) setSequenceNumber(b bool) {

}

func (f *BdaOptFlds) setReportTimestamp(b bool) {

}

func (f *BdaOptFlds) setReasonForInclusion(b bool) {

}

func (f *BdaOptFlds) setDataSetName(b bool) {

}

func (f *BdaOptFlds) setBufferOverflow(b bool) {

}

func (f *BdaOptFlds) setEntryId(b bool) {

}

func NewBdaOptFlds(objectReference *ObjectReference, fc string) *BdaOptFlds {
	b := &BdaOptFlds{}
	b.BdaBitString = *NewBdaBitString(objectReference, fc, "", 10, false, false)
	b.basicType = OPTFLDS
	b.value = []byte{0x02, 0x00}
	return b
}
