package src

import (
	"bytes"
	"strconv"
)

type FileOpenResponse struct {
	tag            *BerTag
	frsmID         *Integer32
	fileAttributes *FileAttributes
	code           []byte
}

func (r *FileOpenResponse) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += r.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		r.frsmID = NewInteger32(0)
		vByteCount += r.frsmID.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 1) {
		r.fileAttributes = NewFileAttributes()
		vByteCount += r.fileAttributes.decode(is, false)
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

func (r *FileOpenResponse) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	codeLength += r.fileAttributes.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	codeLength += r.frsmID.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func NewFileOpenResponse() *FileOpenResponse {
	return &FileOpenResponse{tag: NewBerTag(0, 32, 16)}
}
