package src

type InitRequestDetail struct {
	servicesSupportedCalling *ServiceSupportOptions
	proposedParameterCBB     *ParameterSupportOptions
	proposedVersionNumber    *Integer16
	Tag                      *BerTag
}

func (d *InitRequestDetail) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if (code != nil) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	codeLength := 0
	codeLength += d.servicesSupportedCalling.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.writeByte(0x82)
	codeLength += 1

	codeLength += d.proposedParameterCBB.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	codeLength += d.proposedVersionNumber.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += d.Tag.encode(reverseOS)
	}

	return codeLength

}
func (d *InitRequestDetail) decode() {
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
		proposedVersionNumber = NewInteger16();
		vByteCount += proposedVersionNumber.decode(is, false);
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	if (berTag.equals(128, 0, 1)) {
		proposedParameterCBB = NewParameterSupportOptions();
		vByteCount += proposedParameterCBB.decode(is, false);
		vByteCount += berTag.decode(is);
	} else {
		throw("Tag does not match mandatory sequence component.");
	}

	if (berTag.equals(128, 0, 2)) {
		servicesSupportedCalling = NewServiceSupportOptions();
		vByteCount += servicesSupportedCalling.decode(is, false);
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
		"Unexpected end of sequence, length tag: "
	+ lengthVal
	+ ", bytes decoded: "
	+ vByteCount);
}

func NewInitRequestDetail() *InitRequestDetail {
	return &InitRequestDetail{Tag: NewBerTag(0, 32, 16)}
}
