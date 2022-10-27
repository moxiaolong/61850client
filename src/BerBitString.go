package src

type BerBitString struct {
	value   []byte
	numBits int
	tag     *BerTag
	code    []byte
}

func NewBerBitString(code []byte, value []byte, numBits int) *BerBitString {
	return &BerBitString{tag: NewBerTag(0, 0, 3), numBits: numBits, value: value, code: code}
}

func (o *BerBitString) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if o.code != nil {
		reverseOS.write(o.code)
		if withTag {
			return o.tag.encode(reverseOS) + len(o.code)
		} else {
			return len(o.code)
		}

	} else {

		codeLength := 0
		for codeLength = len(o.value) - 1; codeLength >= 0; codeLength-- {
			reverseOS.writeByte(o.value[codeLength])
		}

		reverseOS.writeByte(byte(len(o.value)*8 - o.numBits))
		codeLength = len(o.value) + 1
		codeLength += encodeLength(reverseOS, codeLength)
		if withTag {
			codeLength += o.tag.encode(reverseOS)
		}

		return codeLength
	}

}

func (o *BerBitString) getValueAsBooleans() interface{} {

}
