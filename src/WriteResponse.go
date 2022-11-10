package src

import (
	"bytes"
	"strconv"
)

type WriteResponse struct {
	tag   *BerTag
	seqOf []*WriteResponseCHOICE
	code  []byte
}

func (r *WriteResponse) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0
	berTag := NewBerTag(0, 0, 0)
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

		element := NewWriteResponseCHOICE()
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

func (r *WriteResponse) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
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

func NewWriteResponse() *WriteResponse {
	return &WriteResponse{tag: NewBerTag(0, 32, 16)}
}
