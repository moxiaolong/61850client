package src

import (
	"bytes"
	"strconv"
)

type DataStructure struct {
	tag   *BerTag
	code  []byte
	seqOf []*Data
}

func NewDataStructure() *DataStructure {
	return &DataStructure{tag: NewBerTag(0, 32, 16)}
}

func (d *DataStructure) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if d.code != nil {
		reverseOS.write(d.code)
		if withTag {
			return d.tag.encode(reverseOS) + len(d.code)
		}
		return len(d.code)
	}
	codeLength := 0
	for _, data := range d.seqOf {
		codeLength += data.encode(reverseOS)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += d.tag.encode(reverseOS)
	}

	return codeLength
}

func (d *DataStructure) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0

	vByteCount := 0

	numDecodedBytes := 0

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

		element := NewData()
		numDecodedBytes = element.decode(is, berTag)
		if numDecodedBytes == 0 {
			throw("Tag did not match")
		}
		vByteCount += numDecodedBytes
		d.seqOf = append(d.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected " + strconv.Itoa(lengthVal) + strconv.Itoa(vByteCount) + " but has ")
	}
	return tlByteCount + vByteCount
}
