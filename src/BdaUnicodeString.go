package src

type BdaUnicodeString struct {
	BasicDataAttribute
	maxLength int
	value     []byte
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
