package src

import "bytes"

type BerGraphicString struct {
	tag *BerTag
	BerOctetString
}

func NewBerGraphicString() *BerGraphicString {
	return &BerGraphicString{tag: NewBerTag(0, 0, 25)}
}

func (b *BerGraphicString) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := b.BerOctetString.encode(reverseOS, false)
	if withTag {
		codeLength += b.tag.encode(reverseOS)
	}

	return codeLength
}

func (b *BerGraphicString) decode(is *bytes.Buffer, withTag bool) int {

	codeLength := 0
	if withTag {
		codeLength += b.tag.decodeAndCheck(is)
	}

	codeLength += b.BerOctetString.decode(is, false)
	return codeLength
}
