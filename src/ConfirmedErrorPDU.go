package src

import (
	"bytes"
	"strconv"
)

type ConfirmedErrorPDU struct {
	invokeID         *Unsigned32
	tag              *BerTag
	code             []byte
	modifierPosition *Unsigned32
	serviceError     *ServiceError
}

func (p *ConfirmedErrorPDU) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += p.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		p.invokeID = NewUnsigned32(0)
		vByteCount += p.invokeID.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		p.modifierPosition = NewUnsigned32(0)
		vByteCount += p.modifierPosition.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 2) {
		p.serviceError = NewServiceError()
		vByteCount += p.serviceError.decode(is, false)
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

func (p *ConfirmedErrorPDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if p.code != nil {
		reverseOS.write(p.code)
		if withTag {
			return p.tag.encode(reverseOS) + len(p.code)
		}
		return len(p.code)
	}

	codeLength := 0
	codeLength += p.serviceError.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 2
	reverseOS.writeByte(0xA2)
	codeLength += 1

	if p.modifierPosition != nil {
		codeLength += p.modifierPosition.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
	}

	codeLength += p.invokeID.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.tag.encode(reverseOS)
	}

	return codeLength
}

func NewConfirmedErrorPDU() *ConfirmedErrorPDU {
	return &ConfirmedErrorPDU{tag: NewBerTag(0, 32, 16)}
}
