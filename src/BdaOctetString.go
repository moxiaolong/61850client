package src

import (
	"strconv"
)

type BdaOctetString struct {
	BasicDataAttribute
	value     []byte
	maxLength int
	mirror    *BdaOctetString
}

func (f *BdaOctetString) getMmsDataObj() *Data {
	data := NewData()
	data.octetString = NewBerOctetString(f.value)
	return data
}

func (f *BdaOctetString) copy() ModelNodeI {
	newCopy := NewBdaOctetString(f.ObjectReference, f.Fc, f.sAddr, f.maxLength, f.dchg, f.dupd)
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

func (s *BdaOctetString) setValueFromMmsDataObj(data *Data) {
	if data.octetString == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: octet_string")
	}
	s.value = data.octetString.value
}

func (s *BdaOctetString) setValue(value []byte) {
	if value != nil && len(value) > s.maxLength {
		throw("OCTET_STRING value size exceeds maxLength of ", strconv.Itoa(s.maxLength))
	}
	s.value = value
}

func (s *BdaOctetString) setDefault() {
	s.value = make([]byte, 0)
}

func NewBdaOctetString(objectReference *ObjectReference, fc string, sAddr string, maxLength int, dchg bool, dupd bool) *BdaOctetString {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	b := &BdaOctetString{BasicDataAttribute: *attribute}
	b.basicType = OCTET_STRING
	b.maxLength = maxLength
	b.setDefault()
	return b
}
