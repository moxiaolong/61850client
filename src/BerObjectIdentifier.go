package src

import (
	"bytes"
	"math"
)

type BerObjectIdentifier struct {
	tag   *BerTag
	code  []byte
	value []int
}

func (t *BerObjectIdentifier) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if t.code != nil {
		reverseOS.write(t.code)
		if withTag {
			return t.tag.encode(reverseOS) + len(t.code)
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
			codeLength += t.tag.encode(reverseOS)
		}

		return codeLength
	}
}
func (t *BerObjectIdentifier) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0
	if withTag {
		codeLength += t.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	codeLength += length.decode(is)
	if length.val == 0 {
		t.value = make([]int, 0)
		return codeLength
	} else {
		byteCode := make([]byte, length.val)
		_, err := is.Read(byteCode)
		if err != nil {
			panic(err)
		}

		codeLength += length.val
		objectIdentifierComponentsList := make([]int, 0)
		subIDEndIndex := 0
		for subIDEndIndex = 0; (byteCode[subIDEndIndex] & 128) == 128; subIDEndIndex++ {
			if subIDEndIndex >= length.val-1 {
				throw("Invalid Object Identifier")
			}
		}

		subidentifier := 0
		i := 0
		for i = 0; i <= subIDEndIndex; i++ {
			subidentifier |= int((byteCode[i]&127)<<(subIDEndIndex-i)) * 7
		}

		if subidentifier < 40 {
			objectIdentifierComponentsList = append(objectIdentifierComponentsList, 0)
			objectIdentifierComponentsList = append(objectIdentifierComponentsList, subidentifier)
		} else if subidentifier < 80 {
			objectIdentifierComponentsList = append(objectIdentifierComponentsList, 1)
			objectIdentifierComponentsList = append(objectIdentifierComponentsList, subidentifier-40)
		} else {
			objectIdentifierComponentsList = append(objectIdentifierComponentsList, 2)
			objectIdentifierComponentsList = append(objectIdentifierComponentsList, subidentifier-80)
		}

		subIDEndIndex++

		for subIDEndIndex < length.val {
			for i = subIDEndIndex; (byteCode[subIDEndIndex] & 128) == 128; subIDEndIndex++ {
				if subIDEndIndex == length.val-1 {
					throw("Invalid Object Identifier")
				}
			}

			subidentifier = 0

			for j := i; j <= subIDEndIndex; j++ {
				subidentifier |= int((byteCode[j]&127)<<(subIDEndIndex-j)) * 7
			}

			objectIdentifierComponentsList = append(objectIdentifierComponentsList, subidentifier)
			subIDEndIndex++
		}

		t.value = objectIdentifierComponentsList

		return codeLength
	}
}

func NewBerObjectIdentifier(code []byte) *BerObjectIdentifier {
	return &BerObjectIdentifier{code: code, tag: NewBerTag(0, 0, 6)}
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

	return &BerObjectIdentifier{value: value, tag: NewBerTag(0, 0, 6)}
}
