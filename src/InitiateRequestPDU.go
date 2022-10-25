package src

import "bytes"

type InitiateRequestPDU struct {
	LocalDetailCalling                *Integer32
	ProposedMaxServOutstandingCalling *Integer16
	ProposedMaxServOutstandingCalled  *Integer16
	ProposedDataStructureNestingLevel *Integer8
	InitRequestDetail                 *InitRequestDetail
	Tag                               *BerTag
}

func (p *InitiateRequestPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func (p *InitiateRequestPDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	codeLength += p.InitRequestDetail.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 4
	reverseOS.writeByte(0xA4)
	codeLength += 1

	if p.ProposedDataStructureNestingLevel != nil {
		codeLength += p.ProposedDataStructureNestingLevel.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
	}

	codeLength += p.ProposedMaxServOutstandingCalled.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.writeByte(0x82)
	codeLength += 1

	codeLength += p.ProposedMaxServOutstandingCalling.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	if p.LocalDetailCalling != nil {
		codeLength += p.LocalDetailCalling.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.Tag.encode(reverseOS)
	}

	return codeLength
}

func NewInitiateRequestPDU() *InitiateRequestPDU {
	return &InitiateRequestPDU{Tag: NewBerTag(0, 32, 16)}
}
