package src

import (
	"bytes"
	"strconv"
)

type ListOfAccessResult struct {
	seqOf []*AccessResult
	tag   *BerTag
	code  []byte
}

func (r *ListOfAccessResult) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()
	numDecodedBytes := 0

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

		element := NewAccessResult()
		numDecodedBytes = element.decode(is, berTag)
		if numDecodedBytes == 0 {
			throw("tag did not match")
		}
		vByteCount += numDecodedBytes
		r.seqOf = append(r.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected " + strconv.Itoa(lengthVal) + " but has " + strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func (r *ListOfAccessResult) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	for i := len(r.seqOf) - 1; i >= 0; i-- {
		codeLength += r.seqOf[i].encode(reverseOS)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func NewListOfAccessResult() *ListOfAccessResult {
	return &ListOfAccessResult{tag: NewBerTag(0, 32, 16)}
}
