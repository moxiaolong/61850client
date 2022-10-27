package src

type ModeSelector struct {
	modeValue *BerInteger
	tag       *BerTag
}

func (s *ModeSelector) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	codeLength += s.modeValue.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func NewModeSelector() *ModeSelector {
	return &ModeSelector{tag: NewBerTag(0, 32, 17)}
}
