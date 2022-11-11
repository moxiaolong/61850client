package src

import (
	"bytes"
	"strconv"
)

type GetNameListRequest struct {
	tag           *BerTag
	objectClass   *ObjectClass
	objectScope   *ObjectScope
	continueAfter *Identifier
	code          []byte
}

func (r *GetNameListRequest) decode(is *bytes.Buffer, withTag bool) int {
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
		r.objectClass = NewObjectClass()
		vByteCount += r.objectClass.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 1) {
		vByteCount += length.decode(is)
		r.objectScope = NewObjectScope()
		vByteCount += r.objectScope.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 2) {
		r.continueAfter = NewIdentifier()
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

func (r *GetNameListRequest) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	sublength := 0

	if r.continueAfter != nil {
		codeLength += r.continueAfter.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
	}

	sublength = r.objectScope.encode(reverseOS)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	sublength = r.objectClass.encode(reverseOS)
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

func NewGetNameListRequest() *GetNameListRequest {
	return &GetNameListRequest{tag: NewBerTag(0, 32, 16)}
}
