package src

type CPTypeNormalModeParameters struct {
	CallingPresentationSelector       *CallingPresentationSelector
	CalledPresentationSelector        *CalledPresentationSelector
	PresentationContextDefinitionList *PresentationContextDefinitionList
	UserData                          *UserData
	tag                               *BerTag
}

func (t *CPTypeNormalModeParameters) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	if (code != nil) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	int codeLength = 0;
	if (userData != nil) {
		codeLength += userData.encode(reverseOS);
	}

	if (userSessionRequirements != nil) {
		codeLength += userSessionRequirements.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.write(0x89);
		codeLength += 1;
	}

	if (presentationRequirements != nil) {
		codeLength += presentationRequirements.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 8
		reverseOS.write(0x88);
		codeLength += 1;
	}

	if (defaultContextName != nil) {
		codeLength += defaultContextName.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 6
		reverseOS.write(0xA6);
		codeLength += 1;
	}

	if (presentationContextDefinitionList != nil) {
		codeLength += presentationContextDefinitionList.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 4
		reverseOS.write(0xA4);
		codeLength += 1;
	}

	if (calledPresentationSelector != nil) {
		codeLength += calledPresentationSelector.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.write(0x82);
		codeLength += 1;
	}

	if (callingPresentationSelector != nil) {
		codeLength += callingPresentationSelector.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.write(0x81);
		codeLength += 1;
	}

	if (protocolVersion != nil) {
		codeLength += protocolVersion.encode(reverseOS, false);
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

func (t *CPTypeNormalModeParameters) decode()  {
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
	if (lengthVal == 0) {
		return tlByteCount;
	}
	vByteCount += berTag.decode(is);

	if (berTag.equals(128, 0, 0)) {
		protocolVersion = NewProtocolVersion();
		vByteCount += protocolVersion.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 0, 1)) {
		callingPresentationSelector = NewCallingPresentationSelector();
		vByteCount += callingPresentationSelector.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 0, 2)) {
		calledPresentationSelector = NewCalledPresentationSelector();
		vByteCount += calledPresentationSelector.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 32, 4)) {
		presentationContextDefinitionList = NewPresentationContextDefinitionList();
		vByteCount += presentationContextDefinitionList.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 32, 6)) {
		defaultContextName = NewDefaultContextName();
		vByteCount += defaultContextName.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 0, 8)) {
		presentationRequirements = NewPresentationRequirements();
		vByteCount += presentationRequirements.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 0, 9)) {
		userSessionRequirements = NewUserSessionRequirements();
		vByteCount += userSessionRequirements.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	}

	userData = NewUserData();
	numDecodedBytes = userData.decode(is, berTag);
	if (numDecodedBytes != 0) {
		vByteCount += numDecodedBytes;
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	} else {
		userData = nil;
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

func NewCPTypeNormalModeParameters() *CPTypeNormalModeParameters {
	return &CPTypeNormalModeParameters{tag: NewBerTag(0, 32, 16)}
}
