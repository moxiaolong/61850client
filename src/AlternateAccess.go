package src

import (
	"bytes"
	"strconv"
)

type AlternateAccess struct {
	seqOf []*AlternateAccessCHOICE
	tag   *BerTag
	code  []byte
}

func (a *AlternateAccess) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if a.code != nil {
		reverseOS.write(a.code)
		if withTag {
			return a.tag.encode(reverseOS) + len(a.code)
		}
		return len(a.code)
	}

	codeLength := 0
	for _, item := range a.seqOf {
		codeLength += item.encode(reverseOS)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += a.tag.encode(reverseOS)
	}

	return codeLength
}

func (a *AlternateAccess) decode(is *bytes.Buffer, withTag bool) int {

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

	for vByteCount < lengthVal || lengthVal < 0 {
		vByteCount += berTag.decode(is)

		if lengthVal < 0 && berTag.equals(0, 0, 0) {
			vByteCount += readEocByte(is)
			break
		}

		element := NewCHOICE()
		numDecodedBytes = element.decode(is, berTag)
		if numDecodedBytes == 0 {
			throw("Tag did not match")
		}
		vByteCount += numDecodedBytes
		a.seqOf = append(a.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected " + strconv.Itoa(lengthVal) + " but has " + strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func NewAlternateAccess() *AlternateAccess {
	return &AlternateAccess{tag: NewBerTag(0, 32, 16)}
}
