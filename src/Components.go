package src

import (
	"bytes"
	"strconv"
)

type TypeDescriptionComponents struct {
	code  []byte
	tag   *BerTag
	seqOf []*TypeDescriptionSEQUENCE
}

func (c *TypeDescriptionComponents) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if c.code != nil {
		reverseOS.write(c.code)
		if withTag {
			return c.tag.encode(reverseOS) + len(c.code)
		}
		return len(c.code)
	}

	codeLength := 0
	for i := len(c.seqOf) - 1; i >= 0; i-- {
		codeLength += c.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += c.tag.encode(reverseOS)
	}

	return codeLength
}

func (c *TypeDescriptionComponents) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()
	if withTag {
		tlByteCount += c.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val

	for vByteCount < lengthVal || lengthVal < 0 {
		vByteCount += berTag.decode(is)

		if lengthVal < 0 && berTag.equals(0, 0, 0) {
			vByteCount += readEocByte(is)
			break
		}

		if !berTag.equals(0, 32, 16) {
			throw("tag does not match mandatory sequence of/set of component.")
		}
		//TODO
		element := NewTypeDescriptionSEQUENCE()
		vByteCount += element.decode(is, false)
		c.seqOf = append(c.seqOf, element)
	}
	if lengthVal >= 0 && vByteCount != lengthVal {
		throw("Decoded SequenceOf or SetOf has wrong length. Expected " + strconv.Itoa(lengthVal) + " but has " + strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func (c *TypeDescriptionComponents) getSEQUENCE() []*TypeDescriptionSEQUENCE {
	if c.seqOf == nil {
		c.seqOf = make([]*TypeDescriptionSEQUENCE, 0)
	}
	return c.seqOf
}

func NewComponents() *TypeDescriptionComponents {
	return &TypeDescriptionComponents{tag: NewBerTag(0, 32, 16)}
}
