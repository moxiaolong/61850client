package src

type BerInteger struct {
	value []byte
	Tag   *BerTag
}

func NewBerInteger(value []byte) *BerInteger {
	return &BerInteger{value: value}
}

func (f *AEQualifierForm2) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	encoded := f.value
	codeLength := len(encoded)
	reverseOS.write(encoded)
	codeLength += encodeLength(reverseOS, codeLength)
	if withTag {
		codeLength += f.Tag.encode(reverseOS)
	}

	return codeLength

}
