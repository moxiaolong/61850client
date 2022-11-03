package src

import "bytes"

type BerOctetString struct {
	tag   *BerTag
	value []byte
}

func NewBerOctetString(value []byte) *BerOctetString {
	return &BerOctetString{tag: NewBerTag(0, 0, 4), value: value}
}

func (b *BerOctetString) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	reverseOS.write(b.value)
	codeLength := len(b.value)
	codeLength += encodeLength(reverseOS, codeLength)
	if withTag {
		codeLength += b.tag.encode(reverseOS)
	}

	return codeLength
}

func (b *BerOctetString) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0
	if withTag {
		codeLength += b.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	codeLength += length.decode(is)
	b.value = make([]byte, length.val)
	if length.val != 0 {
		_, err := is.Read(b.value)
		if err != nil {
			panic(err)
		}
		codeLength += length.val
	}

	return codeLength
}
