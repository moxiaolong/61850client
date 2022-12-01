package src

import (
	"bytes"
	"strconv"
)

type FileDirectoryResponse struct {
	tag                  *BerTag
	listOfDirectoryEntry *ListOfDirectoryEntry
	moreFollows          *BerBoolean
	code                 []byte
}

func (r *FileDirectoryResponse) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += r.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 32, 0) {
		vByteCount += length.decode(is)
		r.listOfDirectoryEntry = NewListOfDirectoryEntry()
		vByteCount += r.listOfDirectoryEntry.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		r.moreFollows = NewBerBoolean(false)
		vByteCount += r.moreFollows.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw(
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(lengthVal))
	return 0
}

func (r *FileDirectoryResponse) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	sublength := 0

	if r.moreFollows != nil {
		codeLength += r.moreFollows.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
	}

	sublength = r.listOfDirectoryEntry.encode(reverseOS, true)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.writeByte(0xA0)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func NewFileDirectoryResponse() *FileDirectoryResponse {
	return &FileDirectoryResponse{tag: NewBerTag(0, 32, 16)}
}
