package src

import (
	"bytes"
	"strconv"
)

type InitResponseDetail struct {
	negotiatedVersionNumber *Integer16
	servicesSupportedCalled *ServiceSupportOptions
	tag                     *BerTag
	negotiatedParameterCBB  *ParameterSupportOptions
	code                    []byte
}

func (d *InitResponseDetail) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += d.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		d.negotiatedVersionNumber = NewInteger16(nil, 0)
		vByteCount += d.negotiatedVersionNumber.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		d.negotiatedParameterCBB = NewParameterSupportOptions(nil)
		vByteCount += d.negotiatedParameterCBB.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 2) {
		d.servicesSupportedCalled = NewServiceSupportOptions(nil, 0)
		vByteCount += d.servicesSupportedCalled.decode(is, false)
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
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (d *InitResponseDetail) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if d.code != nil {
		reverseOS.write(d.code)
		if withTag {
			return d.tag.encode(reverseOS) + len(d.code)
		}
		return len(d.code)
	}

	codeLength := 0
	codeLength += d.servicesSupportedCalled.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.writeByte(0x82)
	codeLength += 1

	codeLength += d.negotiatedParameterCBB.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	codeLength += d.negotiatedVersionNumber.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += d.tag.encode(reverseOS)
	}

	return codeLength
}

func NewInitResponseDetail() *InitResponseDetail {
	return &InitResponseDetail{tag: NewBerTag(0, 32, 16)}
}
