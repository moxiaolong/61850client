package src

import (
	"bytes"
	"strconv"
)

type DomainSpecific struct {
	tag      *BerTag
	domainID *Identifier
	code     []byte
	itemID   *Identifier
}

func (s *DomainSpecific) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += s.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(0, 0, 26) {
		s.domainID = NewIdentifier(nil)
		vByteCount += s.domainID.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(0, 0, 26) {
		s.itemID = NewIdentifier(nil)
		vByteCount += s.itemID.decode(is, false)
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

func (s *DomainSpecific) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	if s.code != nil {
		reverseOS.write(s.code)
		if withTag {
			return s.tag.encode(reverseOS) + len(s.code)
		}
		return len(s.code)
	}

	codeLength := 0
	codeLength += s.itemID.encode(reverseOS, true)

	codeLength += s.domainID.encode(reverseOS, true)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func NewDomainSpecific() *DomainSpecific {
	return &DomainSpecific{tag: NewBerTag(0, 32, 16)}
}
