package src

import (
	"bytes"
	"strconv"
)

type FileName struct {
	tag   *BerTag
	seqOf []*BerGraphicString
	code  []byte
}

func NewFileName() *FileName {
	return &FileName{tag: NewBerTag(0, 32, 16)}
}

func (r *FileName) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0

	vByteCount := 0

	berTag := NewEmptyBerTag()
	if withTag {
		tlByteCount += r.tag.decodeAndCheck(is)
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

		if !berTag.equals(0, 0, 25) {
			throw("tag does not match mandatory sequence of/set of component.")
		}

		element := NewBerGraphicString()
		vByteCount += element.decode(is, false)
		r.seqOf = append(r.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected " + strconv.Itoa(lengthVal) + " but has " + strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func (r *FileName) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	for i := len(r.seqOf) - 1; i >= 0; i-- {
		codeLength += r.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}
