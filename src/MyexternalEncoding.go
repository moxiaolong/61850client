package src

type MyexternalEncoding struct {
	SingleASN1Type *BerAny
}

func (e *MyexternalEncoding) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if (code != nil) {
		reverseOS.write(code);
		return code.length;
	}

	int codeLength = 0;
	int sublength;

	if (arbitrary != nil) {
		codeLength += arbitrary.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.write(0x82);
		codeLength += 1;
		return codeLength;
	}

	if (octetAligned != nil) {
		codeLength += octetAligned.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.write(0x81);
		codeLength += 1;
		return codeLength;
	}

	if (singleASN1Type != nil) {
		sublength = singleASN1Type.encode(reverseOS);
		codeLength += sublength;
		codeLength += BerLength.encodeLength(reverseOS, sublength);
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.write(0xA0);
		codeLength += 1;
		return codeLength;
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.");
}

func (e *MyexternalEncoding) decode()  {
	int tlvByteCount = 0;
	boolean tagWasPassed = (berTag != nil);

	if (berTag == nil) {
		berTag = NewBerTag(0,0,0);
		tlvByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 32, 0)) {
		BerLength length = NewBerLength();
		tlvByteCount += length.decode(is);
		singleASN1Type = NewBerAny();
		tlvByteCount += singleASN1Type.decode(is, nil);
		tlvByteCount += length.readEocIfIndefinite(is);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 1)) {
		octetAligned = NewBerOctetString();
		tlvByteCount += octetAligned.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 2)) {
		arbitrary = NewBerBitString();
		tlvByteCount += arbitrary.decode(is, false);
		return tlvByteCount;
	}

	if (tagWasPassed) {
		return 0;
	}

	throw("Error decoding CHOICE: Tag " + berTag + " matched to no item.");
}
func NewMyexternalEncoding() *MyexternalEncoding {
	return &MyexternalEncoding{}
}
