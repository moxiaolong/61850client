package src

import "math"

type BerObjectIdentifier struct {
	Tag   *BerTag
	code  []byte
	value []int
}

func (t *BerObjectIdentifier) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if t.code != nil {
		reverseOS.write(t.code)
		if withTag {
			return t.Tag.encode(reverseOS) + len(t.code)
		} else {
			return len(t.code)
		}

	} else {
		firstSubidentifier := 40*t.value[0] + t.value[1]
		codeLength := 0

		for i := len(t.value) - 1; i > 0; i-- {
			subidentifier := 0
			if i == 1 {
				subidentifier = firstSubidentifier
			} else {
				subidentifier = t.value[i]
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
			codeLength += t.Tag.encode(reverseOS)
		}

		return codeLength
	}
}

func NewBerObjectIdentifier(code []byte) *BerObjectIdentifier {
	return &BerObjectIdentifier{code: code, Tag: NewBerTag(0, 0, 6)}
}
func NewBerObjectIdentifierOfValue(value []int) *BerObjectIdentifier {
	if len(value) >= 2 && (value[0] != 0 && value[0] != 1 || value[1] <= 39) && value[0] <= 2 {
		var2 := value
		var3 := len(value)

		for var4 := 0; var4 < var3; var4++ {
			objectIdentifierComponent := var2[var4]
			if objectIdentifierComponent < 0 {
				throw("invalid object identifier components")
			}
		}

	} else {
		throw("invalid object identifier components")
	}

	return &BerObjectIdentifier{value: value, Tag: NewBerTag(0, 0, 6)}
}
