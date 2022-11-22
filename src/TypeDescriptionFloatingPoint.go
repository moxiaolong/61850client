package src

import (
	"bytes"
	"strconv"
)

type TypeDescriptionFloatingPoint struct {
	formatWidth   *Unsigned8
	tag           *BerTag
	exponentWidth *Unsigned8
	code          []byte
}

func (p *TypeDescriptionFloatingPoint) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += p.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(0, 0, 2) {
		p.formatWidth = NewUnsigned8()
		vByteCount += p.formatWidth.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(0, 0, 2) {
		p.exponentWidth = NewUnsigned8()
		vByteCount += p.exponentWidth.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
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

func (p *TypeDescriptionFloatingPoint) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if p.code != nil {
		reverseOS.write(p.code)
		if withTag {
			return p.tag.encode(reverseOS) + len(p.code)
		}
		return len(p.code)
	}

	codeLength := 0
	codeLength += p.exponentWidth.encode(reverseOS, true)

	codeLength += p.formatWidth.encode(reverseOS, true)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.tag.encode(reverseOS)
	}

	return codeLength
}

func NewTypeDescriptionFloatingPoint() *TypeDescriptionFloatingPoint {
	return &TypeDescriptionFloatingPoint{tag: NewBerTag(0, 32, 16)}
}
