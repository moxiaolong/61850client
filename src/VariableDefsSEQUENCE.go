package src

import (
	"bytes"
	"strconv"
)

type VariableDefsSEQUENCE struct {
	variableSpecification *VariableSpecification
	alternateAccess       *AlternateAccess
	tag                   *BerTag
	code                  []byte
}

func (s *VariableDefsSEQUENCE) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if s.code != nil {
		reverseOS.write(s.code)
		if withTag {
			return s.tag.encode(reverseOS) + len(s.code)
		}
		return len(s.code)
	}

	codeLength := 0
	if s.alternateAccess != nil {
		codeLength += s.alternateAccess.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 5
		reverseOS.writeByte(0xA5)
		codeLength += 1
	}

	codeLength += s.variableSpecification.encode(reverseOS)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func (s *VariableDefsSEQUENCE) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0

	vByteCount := 0

	numDecodedBytes := 0

	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += s.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	variableSpecification := NewVariableSpecification()
	numDecodedBytes = variableSpecification.decode(is, berTag)
	if numDecodedBytes != 0 {
		vByteCount += numDecodedBytes
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}
	if berTag.equals(128, 32, 5) {
		s.alternateAccess = NewAlternateAccess()
		vByteCount += s.alternateAccess.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
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

func NewVariableDefsSEQUENCE() *VariableDefsSEQUENCE {
	return &VariableDefsSEQUENCE{tag: NewBerTag(0, 32, 16)}
}
