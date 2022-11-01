package src

type BdaBitString struct {
	BasicDataAttribute
	value      []byte
	maxNumBits int
}

func NewBdaBitString(objectReference *ObjectReference, fc string, sAddr string, maxNumBits int, dchg bool, dupd bool) *BdaBitString {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	b := &BdaBitString{BasicDataAttribute: *attribute}
	b.maxNumBits = maxNumBits
	return b
}

func (s *BdaBitString) setDefault() {
	s.value = make([]byte, (s.maxNumBits-1)/8+1)
}
