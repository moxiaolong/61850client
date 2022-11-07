package src

import (
	"bytes"
	"strconv"
)

type PDVList struct {
	presentationContextIdentifier *PresentationContextIdentifier
	presentationDataValues        *PresentationDataValues
	tag                           *BerTag
	transferSyntaxName            *TransferSyntaxName
}

func (l *PDVList) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if code != nil {
		reverseOS.writeByte(code)
		if withTag {
			return tag.encode(reverseOS) + code.length
		}
		return code.length
	}

	codeLength := 0
	codeLength += l.presentationDataValues.encode(reverseOS)

	codeLength += l.presentationContextIdentifier.encode(reverseOS, true)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += l.tag.encode(reverseOS)
	}

	return codeLength
}

func (l *PDVList) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += l.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(0, 0, 6) {
		l.transferSyntaxName = NewTransferSyntaxName()
		vByteCount += l.transferSyntaxName.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(0, 0, 2) {
		l.presentationContextIdentifier = NewPresentationContextIdentifier(nil)
		vByteCount += l.presentationContextIdentifier.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	l.presentationDataValues = NewPresentationDataValues()
	numDecodedBytes = l.presentationDataValues.decode(is, berTag)
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

	throw("Unexpected end of sequence, length tag: ", strconv.Itoa(lengthVal), ", bytes decoded: ", strconv.Itoa(vByteCount))
	return -1
}

func NewPDVList() *PDVList {
	return &PDVList{tag: NewBerTag(0, 32, 16)}
}
