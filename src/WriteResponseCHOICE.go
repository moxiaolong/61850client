package src

import "bytes"

type WriteResponseCHOICE struct {
	failure *DataAccessError
	success *BerNull
	code    []byte
}

func (c *WriteResponseCHOICE) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 0) {
		c.failure = NewDataAccessError()
		tlvByteCount += c.failure.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 1) {
		c.success = NewBerNull()
		tlvByteCount += c.success.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (c *WriteResponseCHOICE) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if c.code != nil {
		reverseOS.write(c.code)
		return len(c.code)
	}

	codeLength := 0
	if c.success != nil {
		codeLength += c.success.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	if c.failure != nil {
		codeLength += c.failure.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func NewWriteResponseCHOICE() *WriteResponseCHOICE {
	return &WriteResponseCHOICE{}
}
