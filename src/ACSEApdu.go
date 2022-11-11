package src

import (
	"bytes"
)

type ACSEApdu struct {
	aarq *AARQApdu
	aare *AAREApdu
	rlrq *RLRQApdu
	rlre *RLREApdu
}

func (a *ACSEApdu) encode(reverseOS *ReverseByteArrayOutputStream) int {
	codeLength := 0
	if a.aarq != nil {
		codeLength += a.aarq.encode(reverseOS, true)
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return -1
}

func (a *ACSEApdu) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewEmptyBerTag()
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(64, 32, 0) {
		a.aarq = NewAARQApdu()
		tlvByteCount += a.aarq.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(64, 32, 1) {
		a.aare = NewAAREApdu()
		tlvByteCount += a.aare.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(64, 32, 2) {
		a.rlrq = NewRLRQApdu()
		tlvByteCount += a.rlrq.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(64, 32, 3) {
		a.rlre = NewRLREApdu()
		tlvByteCount += a.rlre.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag ", berTag.toString(), " matched to no item.")

	return 0
}

func NewACSEApdu() *ACSEApdu {
	return &ACSEApdu{}
}
