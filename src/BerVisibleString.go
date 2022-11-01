package src

import "bytes"

type BerVisibleString struct {
	value []byte
	tag   *BerTag
}

func (s *BerVisibleString) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0
	if withTag {
		codeLength += s.tag.decodeAndCheck(is)
	}
	length := NewBerLength()
	codeLength += length.decode(is)
	s.value = make([]byte, length.val)
	if length.val != 0 {
		readFully(is, s.value)
		codeLength += length.val
	}

	return codeLength
}

func (s *BerVisibleString) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	reverseOS.write(s.value)
	codeLength := len(s.value)
	codeLength += encodeLength(reverseOS, codeLength)
	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func (s *BerVisibleString) toString() string {
	return string(s.value)
}

func NewBerVisibleString() *BerVisibleString {
	return &BerVisibleString{tag: NewBerTag(0, 0, 26)}
}
