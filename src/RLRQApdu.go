package src

import (
	"bytes"
	"strconv"
)

type RLRQApdu struct {
	code            []byte
	tag             *BerTag
	reason          *ReleaseRequestReason
	userInformation *AssociationInformation
}

func (a *RLRQApdu) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if a.code != nil {
		reverseOS.write(a.code)
		if withTag {
			return a.tag.encode(reverseOS) + len(a.code)
		}
		return len(a.code)
	}

	codeLength := 0
	if a.userInformation != nil {
		codeLength += a.userInformation.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 30
		reverseOS.writeByte(0xBE)
		codeLength += 1
	}

	if a.reason != nil {
		codeLength += a.reason.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += a.tag.encode(reverseOS)
	}

	return codeLength
}
func (a *RLRQApdu) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0

	vByteCount := 0

	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += a.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	if lengthVal == 0 {
		return tlByteCount
	}
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		a.reason = NewReleaseRequestReason()
		vByteCount += a.reason.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 30) {
		a.userInformation = NewAssociationInformation()
		vByteCount += a.userInformation.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw(
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func NewRLRQApdu() *RLRQApdu {
	return &RLRQApdu{tag: NewBerTag(64, 32, 2)}
}
