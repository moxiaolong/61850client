package src

type FullyEncodedData struct {
	seqOf []*PDVList
	tag   *BerTag
}

func (d *FullyEncodedData) getPDVList() []*PDVList {
	if d.seqOf == nil {
		d.seqOf = make([]*PDVList, 0)
	}
	return d.seqOf

}

func (d *FullyEncodedData) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	for i := len(d.seqOf) - 1; i >= 0; i-- {
		codeLength += d.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += d.tag.encode(reverseOS)
	}

	return codeLength
}

func NewFullyEncodedData() *FullyEncodedData {
	return &FullyEncodedData{tag: NewBerTag(0, 32, 16)}
}
