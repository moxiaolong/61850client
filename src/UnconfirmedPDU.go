package src

import (
	"bytes"
	"strconv"
)

type UnconfirmedPDU struct {
	Service *UnconfirmedService
	service *UnconfirmedService
	code    []byte
	tag     *BerTag
}

func (p *UnconfirmedPDU) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += p.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	p.service = NewUnconfirmedService()
	numDecodedBytes = p.service.decode(is, berTag)
	if numDecodedBytes != 0 {
		vByteCount += numDecodedBytes
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}
	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw("Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (p *UnconfirmedPDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if p.code != nil {
		reverseOS.write(p.code)
		if withTag {
			return p.tag.encode(reverseOS) + len(p.code)
		}
		return len(p.code)
	}

	codeLength := 0
	codeLength += p.service.encode(reverseOS)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.tag.encode(reverseOS)
	}

	return codeLength
}

func NewUnconfirmedPDU() *UnconfirmedPDU {
	return &UnconfirmedPDU{tag: NewBerTag(0, 32, 16)}
}
