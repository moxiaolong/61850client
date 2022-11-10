package src

import (
	"bytes"
	"strconv"
)

type ReadResponse struct {
	tag                         *BerTag
	variableAccessSpecification *VariableAccessSpecification
	listOfAccessResult          *ListOfAccessResult
	code                        []byte
}

func (r *ReadResponse) decode(is *bytes.Buffer, withTag bool) int {
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

	if berTag.equals(128, 32, 0) {
		vByteCount += length.decode(is)
		r.variableAccessSpecification = NewVariableAccessSpecification()
		vByteCount += r.variableAccessSpecification.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		r.listOfAccessResult = NewListOfAccessResult()
		vByteCount += r.listOfAccessResult.decode(is, false)
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

func (r *ReadResponse) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	sublength := 0
	codeLength += r.listOfAccessResult.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	if r.variableAccessSpecification != nil {
		sublength = r.variableAccessSpecification.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
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

func NewReadResponse() *ReadResponse {
	return &ReadResponse{tag: NewBerTag(0, 32, 16)}
}
