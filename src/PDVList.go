package src

type PDVList struct {
	PresentationContextIdentifier *PresentationContextIdentifier
	PresentationDataValues        *PDVListPresentationDataValues
	tag                           *BerTag
}

func (l *PDVList) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := 0
	codeLength += l.PresentationDataValues.encode(reverseOS)

	codeLength += l.PresentationContextIdentifier.encode(reverseOS, true)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += l.tag.encode(reverseOS)
	}

	return codeLength
}

func NewPDVList() *PDVList {
	return &PDVList{tag: NewBerTag(0, 32, 16)}
}
