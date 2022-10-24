package src

import (
	"bytes"
)

type ACSEApdu struct {
	Aarq *AARQApdu
	Aare *AAREApdu
}

func (a *ACSEApdu) encode(reverseOS *ReverseByteArrayOutputStream) int {
	codeLength := 0
	if a.Aarq != nil {
		codeLength += a.Aarq.encode(reverseOS, true)
		return codeLength
	}

	Throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func (a *ACSEApdu) decode(is *bytes.Buffer) int {

	tlvByteCount := 0

	berTag := NewBerTag(0, 0, 0)
	tlvByteCount += berTag.decode(is)

	return 0

}

func NewACSEApdu() *ACSEApdu {
	return &ACSEApdu{}
}
