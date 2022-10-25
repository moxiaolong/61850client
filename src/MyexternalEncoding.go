package src

type MyexternalEncoding struct {
	SingleASN1Type *BerAny
}

func (e *MyexternalEncoding) encode(reverseOS *ReverseByteArrayOutputStream) int {
	codeLength := 0
	sublength := 0

	if e.SingleASN1Type != nil {
		sublength = e.SingleASN1Type.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	Throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func NewMyexternalEncoding() *MyexternalEncoding {
	return &MyexternalEncoding{}
}
