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
		buffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(buffer, binary.BigEndian, int64(t.value))

		buf := buffer.Bytes()
		codeLength := len(buf)
		reverseOS.write(buf)
		codeLength += encodeLength(reverseOS, codeLength)
		if withTag {
			codeLength += t.Tag.encode(reverseOS)
		}

		return codeLength
	}
}

func (t *BerInteger) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0
	if withTag {
		codeLength += t.Tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	codeLength += length.decode(is)
	if length.val < 1 {
		throw("Decoded length of BerInteger is not correct")
	} else {

		byteCode := make([]byte, length.val)
		readFully(is, byteCode)
		codeLength += length.val
		//TODO
		t.value = int(binary.LittleEndian.Uint64(byteCode))
		return codeLength
	}
	return -1
}

func NewBerInteger(code []byte, value int) *BerInteger {
	return &BerInteger{code: code, value: value, Tag: NewBerTag(0, 0, 2)}
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
