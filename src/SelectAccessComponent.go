package src

import "bytes"

type SelectAccessComponent struct {
	code  []byte
	basic *BasicIdentifier
}

func (c *SelectAccessComponent) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if c.code != nil {
		reverseOS.write(c.code)
		return len(c.code)
	}

	codeLength := 0
	if c.basic != nil {
		codeLength += c.basic.encode(reverseOS, true)
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func (c *SelectAccessComponent) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0

	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(0, 0, 26) {
		c.basic = NewBasicIdentifier(nil)
		tlvByteCount += c.basic.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func NewSelectAccessComponent() *SelectAccessComponent {
	return &SelectAccessComponent{}
}
