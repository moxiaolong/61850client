package src

import (
	"bytes"
	"strconv"
)

type ListOfVariableListName struct {
	tag   *BerTag
	seqOf []*ObjectName
	code  []byte
}

func (n *ListOfVariableListName) decode(is *bytes.Buffer, withTag bool) int {

	numDecodedBytes := 0
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)
	if withTag {
		tlByteCount += n.tag.decodeAndCheck(is)
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

		element := NewObjectName()
		numDecodedBytes = element.decode(is, berTag)
		if numDecodedBytes == 0 {
			throw("Tag did not match")
		}
		vByteCount += numDecodedBytes
		n.seqOf = append(n.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected " + strconv.Itoa(lengthVal) + " but has " + strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func (n *ListOfVariableListName) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if n.code != nil {
		reverseOS.write(n.code)
		if withTag {
			return n.tag.encode(reverseOS) + len(n.code)
		}
		return len(n.code)
	}

	codeLength := 0
	for i := len(n.seqOf) - 1; i >= 0; i-- {
		codeLength += n.seqOf[i].encode(reverseOS)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += n.tag.encode(reverseOS)
	}

	return codeLength
}

func NewListOfVariableListName() *ListOfVariableListName {
	return &ListOfVariableListName{tag: NewBerTag(0, 32, 16)}
}
