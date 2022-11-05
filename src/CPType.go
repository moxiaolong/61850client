package src

type CPType struct {
	ModeSelector         *ModeSelector
	NormalModeParameters *CPTypeNormalModeParameters
	tag                  *BerTag
}

func (receiver *CPType) decode()  {
	int tlByteCount = 0;
	int vByteCount = 0;
	BerTag berTag = NewBerTag(0,0,0);

	if (withTag) {
		tlByteCount += tag.decodeAndCheck(is);
	}

	BerLength length = NewBerLength();
	tlByteCount += length.decode(is);
	int lengthVal = length.val;

	while (vByteCount < lengthVal || lengthVal < 0) {
		vByteCount += berTag.decode(is);
		if (berTag.equals(128, 32, 0)) {
			modeSelector = NewModeSelector();
			vByteCount += modeSelector.decode(is, false);
		} else if (berTag.equals(128, 32, 2)) {
			normalModeParameters = NewNormalModeParameters();
			vByteCount += normalModeParameters.decode(is, false);
		} else if (lengthVal < 0 && berTag.equals(0, 0, 0)) {
			vByteCount += BerLength.readEocByte(is);
			return tlByteCount + vByteCount;
		} else {
			throw("Tag does not match any set component: " + berTag);
		}
	}
	if (vByteCount != lengthVal) {
		throw(
			"Length of set does not match length tag, length tag: "
		+ lengthVal
		+ ", actual set length: "
		+ vByteCount);
	}
	return tlByteCount + vByteCount;
}

func (t *CPType) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	if t.NormalModeParameters != nil {
		codeLength += t.NormalModeParameters.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
	}

	codeLength += t.ModeSelector.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.writeByte(0xA0)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += t.tag.encode(reverseOS)
	}

	return codeLength
}

func NewCPType() *CPType {
	return &CPType{tag: NewBerTag(0, 32, 17)}
}
