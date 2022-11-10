package src

import "bytes"

type BerBoolean struct {
	value bool
	code  []byte
	tag   *BerTag
}

func (b *BerBoolean) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0

	if withTag {
		codeLength += b.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	codeLength += length.decode(is)

	if length.val != 1 {
		throw("Decoded length of BerBoolean is not correct")
	}

	nextByte, err := is.ReadByte()
	if err != nil {
		throw("Unexpected end of input stream.")
	}

	codeLength++
	b.value = nextByte != 0

	return codeLength
}

func (b *BerBoolean) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if b.code != nil {
		reverseOS.write(b.code)
		if withTag {
			return b.tag.encode(reverseOS) + len(b.code)
		}
		return len(b.code)
	}

	codeLength := 1

	if b.value {
		reverseOS.writeByte(0x01)
	} else {
		reverseOS.writeByte(0)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += b.tag.encode(reverseOS)
	}

	return codeLength
}

func NewBerBoolean() *BerBoolean {
	return &BerBoolean{tag: NewBerTag(0, 0, 1)}
}
