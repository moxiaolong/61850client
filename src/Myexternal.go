package src

type Myexternal struct {
	DirectReference   *BerObjectIdentifier
	IndirectReference *BerInteger
	Encoding          *MyexternalEncoding
	Tag               *BerTag
}

func (m *Myexternal) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	codeLength += m.Encoding.encode(reverseOS)

	if m.IndirectReference != nil {
		codeLength += m.IndirectReference.encode(reverseOS, true)
	}

	if m.DirectReference != nil {
		codeLength += m.DirectReference.encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += m.Tag.encode(reverseOS)
	}

	return codeLength
}

func NewMyexternal() *Myexternal {
	return &Myexternal{Tag: NewBerTag(0, 32, 8)}
}
