package src

import (
	"bytes"
	"strconv"
)

type DirectoryEntry struct {
	tag            *BerTag
	fileAttributes *FileAttributes
	fileName       *FileName
	code           []byte
}

func (e *DirectoryEntry) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += e.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 32, 0) {
		e.fileName = NewFileName()
		vByteCount += e.fileName.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 1) {
		e.fileAttributes = NewFileAttributes()
		vByteCount += e.fileAttributes.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw("Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (e *DirectoryEntry) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if e.code != nil {
		reverseOS.write(e.code)
		if withTag {
			return e.tag.encode(reverseOS) + len(e.code)
		}
		return len(e.code)
	}

	codeLength := 0
	codeLength += e.fileAttributes.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	codeLength += e.fileName.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.writeByte(0xA0)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += e.tag.encode(reverseOS)
	}

	return codeLength
}

func NewDirectoryEntry() *DirectoryEntry {
	return &DirectoryEntry{tag: NewBerTag(0, 32, 16)}
}
