package src

import (
	"bytes"
	"strconv"
)

type InitiateRequestPDU struct {
	LocalDetailCalling                *Integer32
	ProposedMaxServOutstandingCalling *Integer16
	ProposedMaxServOutstandingCalled  *Integer16
	ProposedDataStructureNestingLevel *Integer8
	InitRequestDetail                 *InitRequestDetail
	tag                               *BerTag
	localDetailCalling                *Integer32
	proposedMaxServOutstandingCalling *Integer16
	proposedMaxServOutstandingCalled  *Integer16
	proposedDataStructureNestingLevel *Integer8
	initRequestDetail                 *InitRequestDetail
	code                              []byte
}

func (p *InitiateRequestPDU) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += p.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		p.localDetailCalling = NewInteger32(0)
		vByteCount += p.localDetailCalling.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 1) {
		p.proposedMaxServOutstandingCalling = NewInteger16(nil, 0)
		vByteCount += p.proposedMaxServOutstandingCalling.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 2) {
		p.proposedMaxServOutstandingCalled = NewInteger16(nil, 0)
		vByteCount += p.proposedMaxServOutstandingCalled.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 3) {
		p.proposedDataStructureNestingLevel = NewInteger8(0)
		vByteCount += p.proposedDataStructureNestingLevel.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 4) {
		p.initRequestDetail = NewInitRequestDetail()
		vByteCount += p.initRequestDetail.decode(is, false)
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

func (p *InitiateRequestPDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if p.code != nil {
		reverseOS.writeByte(p.code)
		if withTag {
			return p.tag.encode(reverseOS) + len(p.code)
		}
		return len(p.code)
	}

	codeLength := 0
	codeLength += p.InitRequestDetail.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 4
	reverseOS.writeByte(0xA4)
	codeLength += 1

	if p.ProposedDataStructureNestingLevel != nil {
		codeLength += p.ProposedDataStructureNestingLevel.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
	}

	codeLength += p.ProposedMaxServOutstandingCalled.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.writeByte(0x82)
	codeLength += 1

	codeLength += p.ProposedMaxServOutstandingCalling.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	if p.LocalDetailCalling != nil {
		codeLength += p.LocalDetailCalling.encode(reverseOS, false)
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

func NewInitiateRequestPDU() *InitiateRequestPDU {
	return &InitiateRequestPDU{tag: NewBerTag(0, 32, 16)}
}
