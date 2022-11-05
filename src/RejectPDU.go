package src

import "bytes"

type RejectPDU struct {
	OriginalInvokeID *Unsigned32
}

func (p *RejectPDU) decode(is *bytes.Buffer, b bool) int {
	int tlByteCount = 0;
	int vByteCount = 0;
	int numDecodedBytes;
	BerTag berTag = NewBerTag(0,0,0);

	if (withTag) {
		tlByteCount += tag.decodeAndCheck(is);
	}

	BerLength length = NewBerLength();
	tlByteCount += length.decode(is);
	int lengthVal = length.val;
	vByteCount += berTag.decode(is);

	if (berTag.equals(128, 0, 0)) {
		originalInvokeID = NewUnsigned32();
		vByteCount += originalInvokeID.decode(is, false);
		vByteCount += berTag.decode(is);
	}

	rejectReason = NewRejectReason();
	numDecodedBytes = rejectReason.decode(is, berTag);
	if (numDecodedBytes != 0) {
		vByteCount += numDecodedBytes;
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

func (p *RejectPDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	if (code != nil) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	int codeLength = 0;
	codeLength += rejectReason.encode(reverseOS);

	if (originalInvokeID != nil) {
		codeLength += originalInvokeID.encode(reverseOS, false);
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

func NewRejectPDU() *RejectPDU {
	return &RejectPDU{}

}
