package src

import (
	"bytes"
	"strconv"
)

type Structure struct {
	components *TypeDescriptionComponents
	tag        *BerTag
	code       []byte
	packed     *BerBoolean
}

func (s *Structure) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if s.code != nil {
		reverseOS.write(s.code)
		if withTag {
			return s.tag.encode(reverseOS) + len(s.code)
		}
		return len(s.code)
	}

	codeLength := 0
	codeLength += s.components.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	if s.packed != nil {
		codeLength += s.packed.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func (s *Structure) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0

	vByteCount := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += s.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		s.packed = NewBerBoolean()
		vByteCount += s.packed.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		s.components = NewComponents()
		vByteCount += s.components.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw(
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func NewStructure() *Structure {
	return &Structure{tag: NewBerTag(0, 32, 16)}
}
