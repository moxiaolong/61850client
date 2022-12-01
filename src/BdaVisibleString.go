package src

type BdaVisibleString struct {
	BasicDataAttribute
	maxLength int
	value     []byte
	mirror    *BdaVisibleString
}

func (f *BdaVisibleString) getMmsDataObj() *Data {
	data := NewData()
	data.visibleString = NewBerVisibleString(f.value)
	return data
}

func (f *BdaVisibleString) SetValue(value string) {
	f.value = []byte(value)
}

func (f *BdaVisibleString) copy() ModelNodeI {
	newCopy := NewBdaVisibleString(f.ObjectReference, f.Fc, f.sAddr, f.maxLength, f.dchg, f.dupd)
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

func (s *BdaVisibleString) setValueFromMmsDataObj(data *Data) {
	if data.visibleString == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: visible_string")
	}
	s.value = data.visibleString.value
}

func (s *BdaVisibleString) getStringValue() string {
	return string(s.value)
}

func (s *BdaVisibleString) setDefault() {
	s.value = []byte{0}

}

func NewBdaVisibleString(objectReference *ObjectReference, fc string, sAddr string, maxLength int, dchg bool, dupd bool) *BdaVisibleString {

	b := &BdaVisibleString{BasicDataAttribute: *NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd), maxLength: maxLength}
	b.basicType = VISIBLE_STRING
	b.setDefault()
	return b
}

func (s *BdaVisibleString) GetValueString() string {
	return string(s.value)
}
