package src

import (
	"bytes"
	"strconv"
)

type DeleteNamedVariableListRequest struct {
	tag                    *BerTag
	code                   []byte
	scopeOfDelete          *BerInteger
	listOfVariableListName *ListOfVariableListName
	domainName             *DomainName
}

func (r *DeleteNamedVariableListRequest) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

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

	if berTag.equals(128, 0, 0) {
		r.scopeOfDelete = NewBerInteger(nil, 0)
		vByteCount += r.scopeOfDelete.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		r.listOfVariableListName = NewListOfVariableListName()
		vByteCount += r.listOfVariableListName.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 2) {
		vByteCount += length.decode(is)
		r.domainName = NewDomainName()
		vByteCount += r.domainName.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
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
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (r *DeleteNamedVariableListRequest) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	sublength := 0

	if r.domainName != nil {
		sublength = r.domainName.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
	}

	if r.listOfVariableListName != nil {
		codeLength += r.listOfVariableListName.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
	}

	if r.scopeOfDelete != nil {
		codeLength += r.scopeOfDelete.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func NewDeleteNamedVariableListRequest() *DeleteNamedVariableListRequest {
	return &DeleteNamedVariableListRequest{tag: NewBerTag(0, 32, 16)}
}
