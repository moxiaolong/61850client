package src

import (
	"bytes"
	"strconv"
)

type ServiceError struct {
	ErrorClass *ErrorClass
	Tag        *BerTag
}

func NewServiceError() *ServiceError {

	return &ServiceError{Tag: NewBerTag(0, 32, 16)}
}

func (p *InitiateErrorPDU) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := &BerTag{}

	if withTag {
		tlByteCount += p.Tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 32, 0) {
		vByteCount += length.decode(is)
		errorClass := NewErrorClass()
		vByteCount += errorClass.decode(is)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		p.additionalCode = NewBerInteger(nil, 0)
		vByteCount += p.additionalCode.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 2) {
		p.additionalDescription = NewBerVisibleString()
		vByteCount += p.additionalDescription.decode(is, false)
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
		"Unexpected end of sequence, length tag: ", strconv.Itoa(lengthVal),
		", bytes decoded: ", strconv.Itoa(vByteCount))
	return -1
}

func (p *InitiateErrorPDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	sublength := 0

	if p.additionalDescription != nil {
		codeLength += p.additionalDescription.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
	}

	if p.additionalCode != nil {
		codeLength += p.additionalCode.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
	}

	sublength = p.ErrorClass.encode(reverseOS)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.writeByte(0xA0)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.Tag.encode(reverseOS)
	}

	return codeLength
}
