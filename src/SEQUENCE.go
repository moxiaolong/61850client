package src

type SEQUENCE struct {
	transferSyntaxNameList        *TransferSyntaxNameList
	abstractSyntaxName            *AbstractSyntaxName
	presentationContextIdentifier *PresentationContextIdentifier
	tag                           *BerTag
}

func (s *SEQUENCE) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := 0
	codeLength += s.transferSyntaxNameList.encode(reverseOS, true)

	codeLength += s.abstractSyntaxName.encode(reverseOS, true)

	codeLength += s.presentationContextIdentifier.encode(reverseOS, true)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func NewSEQUENCE() *SEQUENCE {
	return &SEQUENCE{tag: NewBerTag(0, 32, 16)}
}
