package src

import (
	"bytes"
	"strconv"
)

type ContextListSEQUENCE struct {
	presentationContextIdentifier *PresentationContextIdentifier
	abstractSyntaxName            *AbstractSyntaxName
	transferSyntaxNameList        *TransferSyntaxNameList
	tag                           *BerTag
}

func (s *ContextListSEQUENCE) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := 0
	codeLength += s.transferSyntaxNameList.encode(reverseOS, true)

	codeLength += s.abstractSyntaxName.encode(reverseOS, true)

	codeLength += s.presentationContextIdentifier.encode(reverseOS, true)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func (s *ContextListSEQUENCE) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += s.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(0, 0, 2) {
		s.presentationContextIdentifier = NewPresentationContextIdentifier(nil, 0)
		vByteCount += s.presentationContextIdentifier.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(0, 0, 6) {
		s.abstractSyntaxName = NewAbstractSyntaxName()
		vByteCount += s.abstractSyntaxName.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
	}

	if berTag.equals(0, 32, 16) {
		s.transferSyntaxNameList = NewTransferSyntaxNameList()
		vByteCount += s.transferSyntaxNameList.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("Tag does not match mandatory sequence component.")
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

func NewSEQUENCE() *ContextListSEQUENCE {
	return &ContextListSEQUENCE{tag: NewBerTag(0, 32, 16)}
}
