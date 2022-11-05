package src

import "bytes"

type ConfirmedResponsePDU struct {
	invokeID *Unsigned32
}

func (p *ConfirmedResponsePDU) decode(is *bytes.Buffer, withTag bool) int {
	  tlByteCount := 0;
	  vByteCount := 0;
	  numDecodedBytes:=0;
	  berTag := NewBerTag(0,0,0);

	if (withTag) {
		tlByteCount += p.tag.decodeAndCheck(is);
	}

	  length := NewBerLength();
	tlByteCount += length.decode(is);
	int lengthVal = length.val;
	vByteCount += berTag.decode(is);

	if (berTag.equals(Unsigned32.tag)) {
		invokeID = NewUnsigned32();
		vByteCount += invokeID.decode(is, false);
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	service = NewConfirmedServiceResponse();
	numDecodedBytes = service.decode(is, berTag);
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

func (p *ConfirmedResponsePDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	if (code != nil) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	int codeLength = 0;
	codeLength += service.encode(reverseOS);

	codeLength += invokeID.encode(reverseOS, true);

	codeLength += BerLength.encodeLength(reverseOS, codeLength);

	if (withTag) {
		codeLength += tag.encode(reverseOS);
	}

	return codeLength;
}

func NewConfirmedResponsePDU() *ConfirmedResponsePDU {
	return &ConfirmedResponsePDU{}

}
