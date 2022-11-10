package src

import (
	"bytes"
	"strconv"
)

type Array struct {
	tag              *BerTag
	packed           *BerBoolean
	numberOfElements *Unsigned32
	elementType      *TypeSpecification
	code             []byte
}

func (a *Array) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += a.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		a.packed = NewBerBoolean()
		vByteCount += a.packed.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 1) {
		a.numberOfElements = NewUnsigned32()
		vByteCount += a.numberOfElements.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 2) {
		vByteCount += length.decode(is)
		a.elementType = NewTypeSpecification()
		vByteCount += a.elementType.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
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

func (a *Array) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if a.code != nil {
		reverseOS.write(a.code)
		if withTag {
			return a.tag.encode(reverseOS) + len(a.code)
		}
		return len(a.code)
	}

	codeLength := 0

	sublength := 0

	sublength = a.elementType.encode(reverseOS)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
	reverseOS.writeByte(0xA2)
	codeLength += 1

	codeLength += a.numberOfElements.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	if a.packed != nil {
		codeLength += a.packed.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += a.tag.encode(reverseOS)
	}

	return codeLength
}

func NewArray() *Array {
	return &Array{tag: NewBerTag(0, 32, 16)}
}
