package src

import (
	"bytes"
	"strconv"
)

type ModeSelector struct {
	modeValue *BerInteger
	tag       *BerTag
}

func (s *ModeSelector) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	codeLength += s.modeValue.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func (s *ModeSelector) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += s.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val

	for vByteCount < lengthVal || lengthVal < 0 {
		vByteCount += berTag.decode(is)
		if berTag.equals(128, 0, 0) {
			s.modeValue = NewBerInteger(nil, 0)
			vByteCount += s.modeValue.decode(is, false)
		} else if lengthVal < 0 && berTag.equals(0, 0, 0) {
			vByteCount += readEocByte(is)
			return tlByteCount + vByteCount
		} else {
			throw("tag does not match any set component: ", berTag.toString())
		}
	}
	if vByteCount != lengthVal {
		throw(
			"Length of set does not match length tag, length tag: ", strconv.Itoa(lengthVal), ", actual set length: ", strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func NewModeSelector() *ModeSelector {
	return &ModeSelector{tag: NewBerTag(0, 32, 17)}
}
