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
	int tlByteCount = 0;
	int vByteCount = 0;
	BerTag berTag = NewBerTag(0,0,0);

	if (withTag) {
		tlByteCount += tag.decodeAndCheck(is);
	}

	BerLength length = NewBerLength();
	tlByteCount += length.decode(is);
	int lengthVal = length.val;
	vByteCount += berTag.decode(is);

	if (berTag.equals(128, 0, 0)) {
		localDetailCalling = NewInteger32();
		vByteCount += localDetailCalling.decode(is, false);
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 0, 1)) {
		proposedMaxServOutstandingCalling = NewInteger16();
		vByteCount += proposedMaxServOutstandingCalling.decode(is, false);
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	if (berTag.equals(128, 0, 2)) {
		proposedMaxServOutstandingCalled = NewInteger16();
		vByteCount += proposedMaxServOutstandingCalled.decode(is, false);
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	if (berTag.equals(128, 0, 3)) {
		proposedDataStructureNestingLevel = NewInteger8();
		vByteCount += proposedDataStructureNestingLevel.decode(is, false);
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 32, 4)) {
		initRequestDetail = NewInitRequestDetail();
		vByteCount += initRequestDetail.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	if (lengthVal < 0) {
		if (!berTag.equals(0, 0, 0)) {
			throw("Decoded sequence has wrong end of contents octets");
		}
		vByteCount += BerLength.readEocByte(is);
		return tlByteCount + vByteCount;
	}

	throw(
		"Unexpected end of sequence, length tag: " + lengthVal + ", bytes decoded: " + vByteCount);
}

func (p *InitiateRequestPDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if code != nil {
		reverseOS.write(code)
		if withTag {
			return tag.encode(reverseOS) + code.length
		}
		return code.length
	}

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
