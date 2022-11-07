package src

import (
	"bytes"
	"strconv"
)

type InitRequestDetail struct {
	servicesSupportedCalling *ServiceSupportOptions
	proposedParameterCBB     *ParameterSupportOptions
	proposedVersionNumber    *Integer16
	tag                      *BerTag
	code                     []byte
}

func (d *InitRequestDetail) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if d.code != nil {
		reverseOS.writeByte(d.code)
		if withTag {
			return d.tag.encode(reverseOS) + len(d.code)
		}
		return len(d.code)
	}

	codeLength := 0
	codeLength += d.servicesSupportedCalling.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.writeByte(0x82)
	codeLength += 1

	codeLength += d.proposedParameterCBB.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	codeLength += d.proposedVersionNumber.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += d.tag.encode(reverseOS)
	}

	return codeLength

}
func (d *InitRequestDetail) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += d.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		d.proposedVersionNumber = NewInteger16(nil, 0)
		vByteCount += d.proposedVersionNumber.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		d.proposedParameterCBB = NewParameterSupportOptions(nil)
		vByteCount += d.proposedParameterCBB.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 2) {
		d.servicesSupportedCalling = NewServiceSupportOptions(nil)
		vByteCount += d.servicesSupportedCalling.decode(is, false)
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

	throw(
		"Unexpected end of sequence, length tag: ", strconv.Itoa(lengthVal), ", bytes decoded: ", strconv.Itoa(vByteCount))
	return 0
}

func NewInitRequestDetail() *InitRequestDetail {
	return &InitRequestDetail{tag: NewBerTag(0, 32, 16)}
}
