package src

import "bytes"

type BerGeneralizedTime struct {
	BerVisibleString
	tag *BerTag
}

func (s *BerGeneralizedTime) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0
	if withTag {
		codeLength += s.tag.decodeAndCheck(is)
	}

	codeLength += s.BerVisibleString.decode(is, false)
	return codeLength
}

func (s *BerGeneralizedTime) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := s.BerVisibleString.encode(reverseOS, false)
	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}
func NewBerGeneralizedTime() *BerGeneralizedTime {

	time := BerGeneralizedTime{tag: NewBerTag(0, 0, 24)}
	time.BerVisibleString = *NewBerVisibleString(nil)
	return &time
}
