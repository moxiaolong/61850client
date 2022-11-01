package src

import (
	"strconv"
)

type BdaOctetString struct {
	BasicDataAttribute
	value     []byte
	maxLength int
}

func (s *BdaOctetString) setValue(value []byte) {
	if value != nil && len(value) > s.maxLength {
		throw("OCTET_STRING value size exceeds maxLength of ", strconv.Itoa(s.maxLength))
	}
	s.value = value
}

func NewBdaOctetString(objectReference *ObjectReference, fc string, sAddr string, maxLength int, dchg bool, dupd bool) *BdaOctetString {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	b := &BdaOctetString{BasicDataAttribute: *attribute}
	b.basicType = OCTET_STRING
	b.maxLength = maxLength
	b.value = make([]byte, 0)
	return b
}
