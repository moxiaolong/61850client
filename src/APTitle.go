package src

import (
	"bytes"
)

type APTitle struct {
	ApTitleForm2 *ApTitleForm2
}

func (t *APTitle) encode(reverseOS *ReverseByteArrayOutputStream) int {

	codeLength := 0
	if t.ApTitleForm2 != nil {
		codeLength += t.ApTitleForm2.encode(reverseOS, true)
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func (t *APTitle) decode(is *bytes.Buffer, t2 interface{}) int {

}

func NewAPTitle() *APTitle {
	return &APTitle{}
}
