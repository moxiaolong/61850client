package src

import (
	"bytes"
	"encoding/binary"
)

type BerInteger struct {
	code  []byte
	value int
	Tag   *BerTag
}

func (t *BerInteger) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if t.code != nil {
		reverseOS.write(t.code)
		if withTag {
			return t.Tag.encode(reverseOS) + len(t.code)
		} else {
			return len(t.code)
		}

	} else {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(t.value))

		codeLength := len(buf)
		reverseOS.write(buf)
		codeLength += encodeLength(reverseOS, codeLength)
		if withTag {
			codeLength += t.Tag.encode(reverseOS)
		}

		return codeLength
	}
}

func (t *BerInteger) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewBerInteger(value []byte) *BerInteger {
	return &BerInteger{code: value, Tag: NewBerTag(0, 0, 2)}
}

func (f *AEQualifierForm2) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	encoded := f.code
	codeLength := len(encoded)
	reverseOS.write(encoded)
	codeLength += encodeLength(reverseOS, codeLength)
	if withTag {
		codeLength += f.Tag.encode(reverseOS)
	}

	return codeLength

}
