package src

type PDVListPresentationDataValues struct {
	SingleASN1Type *BerAny
}

func (v *PDVListPresentationDataValues) encode(reverseOS *ReverseByteArrayOutputStream) int {
	codeLength := 0
	sublength := 0
	if v.SingleASN1Type != nil {
		sublength = v.SingleASN1Type.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func NewPDVListPresentationDataValues() *PDVListPresentationDataValues {
	return &PDVListPresentationDataValues{}
}
