package src

import (
	"bytes"
	"strconv"
)

type Myexternal struct {
	directReference   *BerObjectIdentifier
	indirectReference *BerInteger
	encoding          *MyexternalEncoding
	tag               *BerTag
	code              []byte
}

func (m *Myexternal) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if m.code != nil {
		reverseOS.write(m.code)
		if withTag {
			return m.tag.encode(reverseOS) + len(m.code)
		}
		return len(m.code)
	}
	codeLength := 0
	codeLength += m.encoding.encode(reverseOS)

	if m.indirectReference != nil {
		codeLength += m.indirectReference.encode(reverseOS, true)
	}

	if m.directReference != nil {
		codeLength += m.directReference.encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += m.tag.encode(reverseOS)
	}

	return codeLength
}

func (m *Myexternal) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += m.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(0, 0, 6) {
		m.directReference = NewBerObjectIdentifier(nil)
		vByteCount += m.directReference.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(0, 0, 2) {
		m.indirectReference = NewBerInteger(nil, 0)
		vByteCount += m.indirectReference.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	m.encoding = NewMyexternalEncoding()
	numDecodedBytes = m.encoding.decode(is, berTag)
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

func NewMyexternal() *Myexternal {
	return &Myexternal{tag: NewBerTag(0, 32, 8)}
}
