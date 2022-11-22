package src

import (
	"bytes"
	"strconv"
)

type SelectAccessIndexRange struct {
	tag              *BerTag
	code             []byte
	lowIndex         *Unsigned32
	numberOfElements *Unsigned32
}

func (r *SelectAccessIndexRange) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	codeLength += r.numberOfElements.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	codeLength += r.lowIndex.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func (r *SelectAccessIndexRange) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0

	vByteCount := 0

	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += r.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		r.lowIndex = NewUnsigned32(0)
		vByteCount += r.lowIndex.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		r.numberOfElements = NewUnsigned32(0)
		vByteCount += r.numberOfElements.decode(is, false)
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

func NewSelectAccessIndexRange() *SelectAccessIndexRange {
	return &SelectAccessIndexRange{tag: NewBerTag(0, 32, 16)}
}
