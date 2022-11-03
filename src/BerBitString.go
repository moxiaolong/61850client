package src

import "bytes"

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

func (o *BerBitString) getValueAsBooleans() []bool {
	if o.value == nil {
		return nil
	} else {
		booleans := make([]bool, o.numBits)

		for i := 0; i < o.numBits; i++ {
			booleans[i] = int(o.value[i/8])&1<<7-i%8 > 0
		}

		return booleans
	}
}

func (o *BerBitString) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0
	if withTag {
		codeLength += o.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	codeLength += length.decode(is)
	o.value = make([]byte, length.val-1)
	unusedBits, err := is.ReadByte()
	if err != nil {
		throw("Unexpected end of input stream.")
	} else if unusedBits > 7 {
		throw("Number of unused bits in bit string expected to be less than 8 but is: ", string(unusedBits))
	} else {
		o.numBits = len(o.value)*8 - int(unusedBits)
		if len(o.value) > 0 {
			_, err := is.Read(o.value)
			if err != nil {
				panic(err)
			}
		}

		codeLength += len(o.value) + 1
		return codeLength
	}
	return -1
}
