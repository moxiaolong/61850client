package src

import (
	"bytes"
	"strconv"
)

type FullyEncodedData struct {
	seqOf []*PDVList
	tag   *BerTag
	code  []byte
}

func (d *FullyEncodedData) getPDVList() []*PDVList {
	if d.seqOf == nil {
		d.seqOf = make([]*PDVList, 0)
	}
	return d.seqOf

}

func (d *FullyEncodedData) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if d.code != nil {
		reverseOS.write(d.code)
		if withTag {
			return d.tag.encode(reverseOS) + len(d.code)
		}
		return len(d.code)
	}

	codeLength := 0
	for i := len(d.seqOf) - 1; i >= 0; i-- {
		codeLength += d.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += d.tag.encode(reverseOS)
	}

	return codeLength
}

func (d *FullyEncodedData) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)
	if withTag {
		tlByteCount += d.tag.decodeAndCheck(is)
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

		if !berTag.equals(0, 32, 16) {
			throw("tag does not match mandatory sequence of/set of component.")
		}
		element := NewPDVList()
		vByteCount += element.decode(is, false)
		d.seqOf = append(d.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw("Decoded SequenceOf or SetOf has wrong length. Expected ", strconv.Itoa(lengthVal), " but has ", strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func NewFullyEncodedData() *FullyEncodedData {
	return &FullyEncodedData{tag: NewBerTag(0, 32, 16), seqOf: make([]*PDVList, 0)}
}
