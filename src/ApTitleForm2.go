package src

import "math"

type ApTitleForm2 struct {
	BerObjectIdentifier
}

func (a *ApTitleForm2) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	firstSubidentifier := 40*a.value[0] + a.value[1]
	codeLength := 0

	for i := len(a.value) - 1; i > 0; i-- {
		subidentifier := 0
		if i == 1 {
			subidentifier = firstSubidentifier
		} else {
			subidentifier = a.value[i]
		}

		subIDLength := 1
		for subIDLength = 1; float64(subidentifier) > math.Pow(2.0, (float64(7*subIDLength))-1.0); {
			subIDLength++
		}

		reverseOS.writeByte(byte(subidentifier & 127))

		for j := 1; j <= subIDLength-1; j++ {
			reverseOS.writeByte(byte(subidentifier>>7*j&255 | 128))
		}

		codeLength += subIDLength
	}

	codeLength += encodeLength(reverseOS, codeLength)
	if withTag {
		codeLength += a.tag.encode(reverseOS)
	}

	return codeLength
}

func NewApTitleForm2(value []int) *ApTitleForm2 {
	identifier := NewBerObjectIdentifierOfValue(value)
	return &ApTitleForm2{*identifier}
}
