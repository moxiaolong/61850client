package src

import (
	"bytes"
	"strconv"
)

type CPType struct {
	modeSelector         *ModeSelector
	normalModeParameters *CPTypeNormalModeParameters
	tag                  *BerTag
	code                 []byte
}

func (t *CPType) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += t.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val

	for vByteCount < lengthVal || lengthVal < 0 {
		vByteCount += berTag.decode(is)
		if berTag.equals(128, 32, 0) {
			t.modeSelector = NewModeSelector()
			vByteCount += t.modeSelector.decode(is, false)
		} else if berTag.equals(128, 32, 2) {
			t.normalModeParameters = NewCPTypeNormalModeParameters()
			vByteCount += t.normalModeParameters.decode(is, false)
		} else if lengthVal < 0 && berTag.equals(0, 0, 0) {
			vByteCount += readEocByte(is)
			return tlByteCount + vByteCount
		} else {
			throw("tag does not match any set component: " + berTag.toString())
		}
	}
	if vByteCount != lengthVal {
		throw("Length of set does not match length tag, length tag: ", strconv.Itoa(lengthVal), ", actual set length: ", strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func (t *CPType) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if t.code != nil {
		reverseOS.write(t.code)
		if withTag {
			return t.tag.encode(reverseOS) + len(t.code)
		}
		return len(t.code)
	}

	codeLength := 0
	if t.normalModeParameters != nil {
		codeLength += t.normalModeParameters.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
	}

	codeLength += t.modeSelector.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.writeByte(0xA0)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += t.tag.encode(reverseOS)
	}

	return codeLength
}

func NewCPType() *CPType {
	return &CPType{tag: NewBerTag(0, 32, 17)}
}
