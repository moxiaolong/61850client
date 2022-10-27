package src

type CPTypeNormalModeParameters struct {
	CallingPresentationSelector       *CallingPresentationSelector
	CalledPresentationSelector        *CalledPresentationSelector
	PresentationContextDefinitionList *PresentationContextDefinitionList
	UserData                          *UserData
	tag                               *BerTag
}

func (t *CPTypeNormalModeParameters) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	if t.UserData != nil {
		codeLength += t.UserData.encode(reverseOS)
	}

	if t.PresentationContextDefinitionList != nil {
		codeLength += t.PresentationContextDefinitionList.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 4
		reverseOS.writeByte(0xA4)
		codeLength += 1
	}

	if t.CalledPresentationSelector != nil {
		codeLength += t.CalledPresentationSelector.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
	}

	if t.CallingPresentationSelector != nil {
		codeLength += t.CallingPresentationSelector.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += t.tag.encode(reverseOS)
	}

	return codeLength
}

func NewCPTypeNormalModeParameters() *CPTypeNormalModeParameters {
	return &CPTypeNormalModeParameters{tag: NewBerTag(0, 32, 16)}
}
