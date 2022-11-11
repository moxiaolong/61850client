package src

import (
	"bytes"
	"strconv"
)

type ListOfDirectoryEntry struct {
	tag   *BerTag
	seqOf []*DirectoryEntry
	code  []byte
}

func (e *ListOfDirectoryEntry) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()
	if withTag {
		tlByteCount += e.tag.decodeAndCheck(is)
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

		element := NewDirectoryEntry()
		vByteCount += element.decode(is, false)
		e.seqOf = append(e.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected " + strconv.Itoa(lengthVal) + " but has " + strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func (e *ListOfDirectoryEntry) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if e.code != nil {
		reverseOS.write(e.code)
		if withTag {
			return e.tag.encode(reverseOS) + len(e.code)
		}
		return len(e.code)
	}

	codeLength := 0
	for i := len(e.seqOf) - 1; i >= 0; i-- {
		codeLength += e.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += e.tag.encode(reverseOS)
	}

	return codeLength
}

func NewListOfDirectoryEntry() *ListOfDirectoryEntry {
	return &ListOfDirectoryEntry{tag: NewBerTag(0, 32, 16)}
}
