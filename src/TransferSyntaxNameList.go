package src

type TransferSyntaxNameList struct {
	tag   *BerTag
	seqOf []*TransferSyntaxName
}

func (l *TransferSyntaxNameList) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := 0
	for i := len(l.seqOf) - 1; i >= 0; i-- {
		codeLength += l.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += l.tag.encode(reverseOS)
	}

	return codeLength
}

func NewTransferSyntaxNameList() *TransferSyntaxNameList {
	return &TransferSyntaxNameList{tag: NewBerTag(0, 32, 16)}
}
