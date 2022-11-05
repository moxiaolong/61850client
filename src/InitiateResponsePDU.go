package src

import "bytes"

type InitiateResponsePDU struct {
	LocalDetailCalled                   *LocalDetailCalled
	NegotiatedMaxServOutstandingCalling *NegotiatedMaxServOutstandingCalling
	NegotiatedMaxServOutstandingCalled  *NegotiatedMaxServOutstandingCalled
	NegotiatedDataStructureNestingLevel *NegotiatedDataStructureNestingLevel
	InitResponseDetail                  *InitResponseDetail
}

func (p *InitiateResponsePDU) decode(is *bytes.Buffer, b bool) int {
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
		localDetailCalled = NewInteger32();
		vByteCount += localDetailCalled.decode(is, false);
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 0, 1)) {
		negotiatedMaxServOutstandingCalling = NewInteger16();
		vByteCount += negotiatedMaxServOutstandingCalling.decode(is, false);
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	if (berTag.equals(128, 0, 2)) {
		negotiatedMaxServOutstandingCalled = NewInteger16();
		vByteCount += negotiatedMaxServOutstandingCalled.decode(is, false);
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	if (berTag.equals(128, 0, 3)) {
		negotiatedDataStructureNestingLevel = NewInteger8();
		vByteCount += negotiatedDataStructureNestingLevel.decode(is, false);
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 32, 4)) {
		initResponseDetail = NewInitResponseDetail();
		vByteCount += initResponseDetail.decode(is, false);
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

func (p *InitiateResponsePDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	if (code != nil) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	int codeLength = 0;
	codeLength += initResponseDetail.encode(reverseOS, false);
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 4
	reverseOS.write(0xA4);
	codeLength += 1;

	if (negotiatedDataStructureNestingLevel != nil) {
		codeLength += negotiatedDataStructureNestingLevel.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.write(0x83);
		codeLength += 1;
	}

	codeLength += negotiatedMaxServOutstandingCalled.encode(reverseOS, false);
	// write tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.write(0x82);
	codeLength += 1;

	codeLength += negotiatedMaxServOutstandingCalling.encode(reverseOS, false);
	// write tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.write(0x81);
	codeLength += 1;

	if (localDetailCalled != nil) {
		codeLength += localDetailCalled.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.write(0x80);
		codeLength += 1;
	}

	codeLength += BerLength.encodeLength(reverseOS, codeLength);

	if (withTag) {
		codeLength += tag.encode(reverseOS);
	}

	return codeLength;
}

func NewInitiateResponsePDU() *InitiateResponsePDU {
	return &InitiateResponsePDU{}
}
