package src

import (
	"bytes"
	"strconv"
)

type ResultList struct {
	tag   *BerTag
	seqOf []*SEQUENCE
}

func NewResultList() *ResultList {
	return &ResultList{tag: NewBerTag(0, 32, 16), seqOf: make([]*SEQUENCE, 0)}
}

func (r *ResultList) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
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

		if !berTag.equals(0, 32, 16) {
			throw("Tag does not match mandatory sequence of/set of component.")
		}

		element := NewSEQUENCE()
		vByteCount += element.decode(is, false)
		r.seqOf = append(r.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected ", strconv.Itoa(lengthVal), " but has ", strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}
