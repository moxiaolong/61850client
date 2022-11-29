package src

import "bytes"

type AlternateAccessCHOICE struct {
	unnamed *AlternateAccessSelection
	code    []byte
}

func (c *AlternateAccessCHOICE) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if c.code != nil {
		reverseOS.write(c.code)
		return len(c.code)
	}

	codeLength := 0
	if c.unnamed != nil {
		codeLength += c.unnamed.encode(reverseOS)
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func (c *AlternateAccessCHOICE) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0

	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	numDecodedBytes := 0

	c.unnamed = NewAlternateAccessSelection()
	numDecodedBytes = c.unnamed.decode(is, berTag)
	if numDecodedBytes != 0 {
		return tlvByteCount + numDecodedBytes
	} else {
		c.unnamed = nil
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func NewAlternateAccessCHOICE() *AlternateAccessCHOICE {
	return &AlternateAccessCHOICE{}
}
