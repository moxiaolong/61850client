package src

import (
	"bytes"
	"strconv"
)

type InitiateResponsePDU struct {
	localDetailCalled                   *Integer32
	negotiatedMaxServOutstandingCalling *Integer16
	negotiatedMaxServOutstandingCalled  *Integer16
	negotiatedDataStructureNestingLevel *Integer8
	initResponseDetail                  *InitResponseDetail
	tag                                 *BerTag
	code                                []byte
}

func (p *InitiateResponsePDU) decode(is *bytes.Buffer, withTag bool) int {
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
		p.localDetailCalled = NewInteger32(0)
		vByteCount += p.localDetailCalled.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 1) {
		p.negotiatedMaxServOutstandingCalling = NewInteger16(nil, 0)
		vByteCount += p.negotiatedMaxServOutstandingCalling.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 2) {
		p.negotiatedMaxServOutstandingCalled = NewInteger16(nil, 0)
		vByteCount += p.negotiatedMaxServOutstandingCalled.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 3) {
		p.negotiatedDataStructureNestingLevel = NewInteger8(0)
		vByteCount += p.negotiatedDataStructureNestingLevel.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 4) {
		p.initResponseDetail = NewInitResponseDetail()
		vByteCount += p.initResponseDetail.decode(is, false)
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

	throw("Unexpected end of sequence, length tag: ", strconv.Itoa(lengthVal), ", bytes decoded: ", strconv.Itoa(vByteCount))
	return 0
}

func (p *InitiateResponsePDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if p.code != nil {
		reverseOS.write(p.code)
		if withTag {
			return p.tag.encode(reverseOS) + len(p.code)
		}
		return len(p.code)
	}

	codeLength := 0
	codeLength += p.initResponseDetail.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 4
	reverseOS.writeByte(0xA4)
	codeLength += 1

	if p.negotiatedDataStructureNestingLevel != nil {
		codeLength += p.negotiatedDataStructureNestingLevel.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
	}

	codeLength += p.negotiatedMaxServOutstandingCalled.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.writeByte(0x82)
	codeLength += 1

	codeLength += p.negotiatedMaxServOutstandingCalling.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	if p.localDetailCalled != nil {
		codeLength += p.localDetailCalled.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.tag.encode(reverseOS)
	}

	return codeLength
}

func NewInitiateResponsePDU() *InitiateResponsePDU {
	return &InitiateResponsePDU{tag: NewBerTag(0, 32, 16)}
}
