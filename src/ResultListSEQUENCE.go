package src

import (
	"bytes"
	"strconv"
)

type ResultListSEQUENCE struct {
	transferSyntaxNameList        *TransferSyntaxNameList
	abstractSyntaxName            *AbstractSyntaxName
	presentationContextIdentifier *PresentationContextIdentifier
	tag                           *BerTag
	ComponentName                 *Identifier
	ComponentType                 *TypeSpecification
	result                        *Result
	transferSyntaxName            *TransferSyntaxName
	providerReason                *BerInteger
}

func (s *ResultListSEQUENCE) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
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

func (s *ResultListSEQUENCE) decode(is *bytes.Buffer, withTag bool) int {
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

	if berTag.equals(128, 0, 0) {
		s.result = NewResult()
		vByteCount += s.result.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		s.transferSyntaxName = NewTransferSyntaxName()
		vByteCount += s.transferSyntaxName.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 2) {
		s.providerReason = NewBerInteger(nil, 0)
		vByteCount += s.providerReason.decode(is, false)
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
		"Unexpected end of sequence, length tag: ", strconv.Itoa(lengthVal), ", bytes decoded: ", strconv.Itoa(vByteCount))
	return 0
}

func NewResultListSEQUENCE() *ResultListSEQUENCE {
	return &ResultListSEQUENCE{tag: NewBerTag(0, 32, 16)}
}
