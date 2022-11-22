package src

import (
	"bytes"
	"strconv"
)

type SelectAlternateAccess struct {
	accessSelection *AccessSelection
	alternateAccess *AlternateAccess
	code            []byte
	tag             *BerTag
}

func (a *SelectAlternateAccess) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if a.code != nil {
		reverseOS.write(a.code)
		if withTag {
			return a.tag.encode(reverseOS) + len(a.code)
		}
		return len(a.code)
	}

	codeLength := 0
	codeLength += a.alternateAccess.encode(reverseOS, true)

	codeLength += a.accessSelection.encode(reverseOS)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += a.tag.encode(reverseOS)
	}

	return codeLength
}

func (a *SelectAlternateAccess) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0

	vByteCount := 0

	numDecodedBytes := 0

	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += a.tag.decodeAndCheck(is)
	}
	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	a.accessSelection = NewAccessSelection()
	numDecodedBytes = a.accessSelection.decode(is, berTag)
	if numDecodedBytes != 0 {
		vByteCount += numDecodedBytes
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}
	if berTag.equals(0, 32, 16) {
		a.alternateAccess = NewAlternateAccess()
		vByteCount += a.alternateAccess.decode(is, false)
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

func NewSelectAlternateAccess() *SelectAlternateAccess {
	return &SelectAlternateAccess{}
}
