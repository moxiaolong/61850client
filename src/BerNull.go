package src

import "bytes"

type BerNull struct {
	Tag *BerTag
}

func NewBerNull() *BerNull {
	return &BerNull{Tag: NewBerTag(0, 0, 5)}
}

func (b *BerNull) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := encodeLength(reverseOS, 0)
	if withTag {
		codeLength += b.Tag.encode(reverseOS)
	}

	return codeLength
}

func (b *BerNull) decode(reverseOS *bytes.Buffer, withTag bool) int {

	codeLength := 0
	if withTag {
		codeLength += b.Tag.decodeAndCheck(reverseOS)
	}

	length := NewBerLength()
	codeLength += length.decode(reverseOS)
	if length.val != 0 {
		Throw("Decoded length of BerNull is not correct")
		return -1
	} else {
		return codeLength
	}

}
