package src

import (
	"bytes"
)

type APTitle struct {
	apTitleForm2 *ApTitleForm2
}

func (t *APTitle) encode(reverseOS *ReverseByteArrayOutputStream) int {

	codeLength := 0
	if t.apTitleForm2 != nil {
		codeLength += t.apTitleForm2.encode(reverseOS, true)
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return -1
}

func (t *APTitle) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(0, 0, 6) {
		t.apTitleForm2 = NewApTitleForm2(nil)
		tlvByteCount += t.apTitleForm2.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func NewAPTitle() *APTitle {
	return &APTitle{}
}
