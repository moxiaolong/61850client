package src

import (
	"bytes"
	"math"
	"strconv"
)

type BerTag struct {
	tagClass        int
	primitive       int
	tagNumber       int
	identifierClass int
	tagBytes        []byte
}

func (t *BerTag) decode(is *bytes.Buffer) int {
	nextByte, err := is.ReadByte()

	if err != nil {
		Throw("Unexpected end of input stream.")
	} else {
		t.tagClass = int(nextByte & 192)
		t.primitive = int(nextByte & 32)
		t.tagNumber = int(nextByte & 31)
		codeLength := 1

		if t.tagNumber == 31 {
			t.tagNumber = 0
			numTagBytes := 0
			//do while
			i := 0
			for i == 0 || (nextByte&128) != 0 {
				i++

				nextByte, err := is.ReadByte()
				if err != nil {
					Throw("Unexpected end of input stream.")
				}
				codeLength++
				if numTagBytes >= 6 {
					Throw("Tag is too large.")
				}
				t.tagNumber <<= 7
				t.tagNumber |= int(nextByte & 127)
				numTagBytes++
			}
		}
		return codeLength
	}
	return -1
}

func (t *BerTag) equals(identifierClass int, primitive int, tagNumber int) bool {
	return t.tagNumber == tagNumber && t.tagClass == identifierClass && t.primitive == primitive
}

func (t *BerTag) toString() string {
	return "identifier class: " + strconv.Itoa(t.tagClass) + ", primitive: " + strconv.Itoa(t.primitive) + ", tag number: " + strconv.Itoa(t.tagNumber)

}

func (t *BerTag) code() {
	if t.tagNumber < 31 {
		t.tagBytes = make([]byte, 1)
		t.tagBytes[0] = (byte)(t.tagClass | t.primitive | t.tagNumber)
	} else {

		tagLength := 1
		for float64(t.tagNumber) > math.Pow(2.0, float64(7*tagLength))-1.0 {
			tagLength++
		}

		t.tagBytes = make([]byte, 1+tagLength)
		t.tagBytes[0] = (byte)(t.tagClass | t.primitive | 31)

		for j := 0; j <= tagLength-1; j++ {
			t.tagBytes[j] = (byte)(t.tagNumber>>7*(tagLength-j)&255 | 128)
		}
		t.tagBytes[tagLength] = (byte)(t.tagNumber & 127)
	}
}

func (t *BerTag) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if t.tagBytes == nil {
		t.code()
	}

	for i := len(t.tagBytes) - 1; i >= 0; i-- {
		reverseOS.writeByte(t.tagBytes[i])
	}

	return len(t.tagBytes)
}

func (t *BerTag) decodeAndCheck(reverseOS *bytes.Buffer) int {
	var2 := t.tagBytes
	var3 := len(var2)

	for var4 := 0; var4 < var3; var4++ {
		identifierByte := var2[var4]
		nextByte, err := reverseOS.ReadByte()

		if err != nil {
			Throw("Unexpected end of input stream.")
		}

		if nextByte != (identifierByte & 255) {
			Throw("Identifier does not match, expected: ", string(identifierByte), ", received: ", string(nextByte))
		}
	}

	return len(t.tagBytes)
}

func NewBerTag(identifierClass int, primitive int, tagNumber int) *BerTag {
	b := &BerTag{identifierClass: identifierClass, primitive: primitive, tagNumber: tagNumber}
	b.code()
	return b
}
