package src

import "math"

type ApTitleForm2 struct {
	value []byte
	Tag   *BerTag
}

func (a *ApTitleForm2) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	firstSubidentifier := 40*a.value[0] + a.value[1]
	codeLength := 0

	for i := len(a.value) - 1; i > 0; i-- {
		subidentifier := 0
		if i == 1 {
			subidentifier = int(firstSubidentifier)
		} else {
			subidentifier = int(a.value[i])
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
		codeLength += a.Tag.encode(reverseOS)
	}

	return codeLength
}

func NewApTitleForm2([]int) *ApTitleForm2 {
	return &ApTitleForm2{Tag: NewBerTag(0, 0, 6)}
}
