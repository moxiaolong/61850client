package src

import "bytes"

type ObjectClass struct {
	code             []byte
	basicObjectClass *BerInteger
}

func (c *ObjectClass) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 0) {
		c.basicObjectClass = NewBerInteger(nil, 0)
		tlvByteCount += c.basicObjectClass.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (c *ObjectClass) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if c.code != nil {
		reverseOS.write(c.code)
		return len(c.code)
	}

	codeLength := 0
	if c.basicObjectClass != nil {
		codeLength += c.basicObjectClass.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewObjectClass() *ObjectClass {
	return &ObjectClass{}
}
