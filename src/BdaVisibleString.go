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

func NewBdaVisibleString() *BdaVisibleString {
	return &BdaVisibleString{BasicDataAttribute: *NewBasicDataAttribute(nil, "", "", false, false)}
}
