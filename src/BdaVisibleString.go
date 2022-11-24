package src

type BdaVisibleStringI interface {
}
type BdaVisibleString struct {
	BasicDataAttribute
	maxLength int
	value     []byte
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
