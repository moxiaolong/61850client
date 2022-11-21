package src

type BdaVisibleStringI interface {
}
type BdaVisibleString struct {
	BasicDataAttribute
	value []byte
}

func (s *BdaVisibleString) getStringValue() string {
	return string(s.value)
}

func NewBdaVisibleString(objectReference *ObjectReference, fc string, sAddr string, maxLength int, dchg bool, dupd bool) *BdaVisibleString {
	return &BdaVisibleString{BasicDataAttribute: *NewBasicDataAttribute(nil, "", "", false, false)}
}
