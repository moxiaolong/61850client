package src

import (
	"bytes"
	"strconv"
)

type FileDirectoryRequest struct {
	tag               *BerTag
	fileSpecification *FileName
	continueAfter     *FileName
	code              []byte
}

func (r *FileDirectoryRequest) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += r.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	if lengthVal == 0 {
		return tlByteCount
	}
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 32, 0) {
		r.fileSpecification = NewFileName()
		vByteCount += r.fileSpecification.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		r.continueAfter = NewFileName()
		vByteCount += r.continueAfter.decode(is, false)
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

	throw("Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (r *FileDirectoryRequest) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	if r.continueAfter != nil {
		codeLength += r.continueAfter.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
	}

	if r.fileSpecification != nil {
		codeLength += r.fileSpecification.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength

}

func NewFileDirectoryRequest() *FileDirectoryRequest {
	return &FileDirectoryRequest{tag: NewBerTag(0, 32, 16)}
}
