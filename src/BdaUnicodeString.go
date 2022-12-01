package src

type BdaUnicodeString struct {
	BasicDataAttribute
	maxLength int
	value     []byte
	mirror    *BdaUnicodeString
}

func (f *BdaUnicodeString) getMmsDataObj() *Data {
	data := NewData()
	data.mMSString = NewMMSString(f.value)
	return data

}

func (f *BdaUnicodeString) copy() ModelNodeI {
	newCopy := NewBdaUnicodeString(f.ObjectReference, f.Fc, f.sAddr, f.maxLength, f.dchg, f.dupd)
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

func (s *BdaUnicodeString) setValueFromMmsDataObj(data *Data) {
	if data.mMSString == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: mms_string/unicode_string")
	}
	s.value = data.mMSString.value
}

func (s *BdaUnicodeString) setDefault() {
	s.value = make([]byte, 0)
}

func NewBdaUnicodeString(objectReference *ObjectReference, fc string, sAddr string, maxlenght int, dchg bool, dupd bool) *BdaUnicodeString {

	b := &BdaUnicodeString{}
	b.BasicDataAttribute = *NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	b.basicType = UNICODE_STRING
	b.maxLength = maxlenght
	b.setDefault()

	return b
}
