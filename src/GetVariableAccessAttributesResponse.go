package src

import (
	"bytes"
	"strconv"
)

type GetVariableAccessAttributesResponse struct {
	typeDescription *TypeDescription
	tag             *BerTag
	mmsDeletable    *BerBoolean
	code            []byte
}

func (r *GetVariableAccessAttributesResponse) decode(is *bytes.Buffer, withTag bool) int {
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

	if berTag.equals(128, 0, 0) {
		r.mmsDeletable = NewBerBoolean(false)
		vByteCount += r.mmsDeletable.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 2) {
		vByteCount += length.decode(is)
		r.typeDescription = NewTypeDescription()
		vByteCount += r.typeDescription.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
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

func (r *GetVariableAccessAttributesResponse) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	sublength := 0
	sublength = r.typeDescription.encode(reverseOS)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
	reverseOS.writeByte(0xA2)
	codeLength += 1

	codeLength += r.mmsDeletable.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func NewGetVariableAccessAttributesResponse() *GetVariableAccessAttributesResponse {
	return &GetVariableAccessAttributesResponse{tag: NewBerTag(0, 32, 16)}
}
