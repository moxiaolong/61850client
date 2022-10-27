package src

type CPType struct {
	ModeSelector         *ModeSelector
	NormalModeParameters *CPTypeNormalModeParameters
	tag                  *BerTag
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
