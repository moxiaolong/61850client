package src

import (
	"bytes"
)

type AEQualifier struct {
	aeQualifierForm2 *AEQualifierForm2
}

func (q *AEQualifier) encode(reverseOS *ReverseByteArrayOutputStream) int {

	codeLength := 0
	if q.aeQualifierForm2 != nil {
		codeLength += q.aeQualifierForm2.encode(reverseOS, true)
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return -1
}

func (q *AEQualifier) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewEmptyBerTag()
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(0, 0, 2) {
		q.aeQualifierForm2 = NewAEQualifierForm2(0)
		tlvByteCount += q.aeQualifierForm2.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag ", berTag.toString(), " matched to no item.")
	return 0
}

func NewAEQualifier() *AEQualifier {
	return &AEQualifier{}
}
