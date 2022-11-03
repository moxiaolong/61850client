package src

import (
	"bytes"
)

type AssociationInformation struct {
	seqOf []*Myexternal
	Tag   *BerTag
}

func (a *AssociationInformation) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	codeLength := 0
	for i := len(a.seqOf) - 1; i >= 0; i-- {
		codeLength += a.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += a.Tag.encode(reverseOS)
	}

	return codeLength
}

func (a *AssociationInformation) decode(is *bytes.Buffer, b bool) int {

}

func NewAssociationInformation() *AssociationInformation {
	return &AssociationInformation{Tag: NewBerTag(0, 32, 16)}
}
