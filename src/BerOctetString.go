package src

type BerOctetString struct {
	tag   *BerTag
	value []byte
}

func NewBerOctetString(value []byte) *BerOctetString {
	return &BerOctetString{tag: NewBerTag(0, 0, 4), value: value}
}

func (b *BerOctetString) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	reverseOS.write(b.value)
	codeLength := len(b.value)
	codeLength += encodeLength(reverseOS, codeLength)
	if withTag {
		codeLength += b.tag.encode(reverseOS)
	}

	return codeLength
}
