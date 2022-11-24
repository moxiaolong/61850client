package src

import (
	"bytes"
	"strconv"
)

type TypeDescriptionSEQUENCE struct {
	transferSyntaxNameList        *TransferSyntaxNameList
	abstractSyntaxName            *AbstractSyntaxName
	presentationContextIdentifier *PresentationContextIdentifier
	tag                           *BerTag
	componentName                 *Identifier
	componentType                 *TypeSpecification
	result                        *Result
	transferSyntaxName            *TransferSyntaxName
	providerReason                *BerInteger
	code                          []byte
}

func (s *TypeDescriptionSEQUENCE) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if s.code != nil {
		reverseOS.write(s.code)
		if withTag {
			return s.tag.encode(reverseOS) + len(s.code)
		}
		return len(s.code)
	}

	codeLength := 0
	sublength := 0

	sublength = s.componentType.encode(reverseOS)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	if s.componentName != nil {
		codeLength += s.componentName.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += s.tag.encode(reverseOS)
	}

	return codeLength
}

func (s *TypeDescriptionSEQUENCE) decode(is *bytes.Buffer, withTag bool) int {

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
		s.componentName = NewIdentifier(nil)
		vByteCount += s.componentName.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		vByteCount += length.decode(is)
		s.componentType = NewTypeSpecification()
		vByteCount += s.componentType.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
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

func NewTypeDescriptionSEQUENCE() *TypeDescriptionSEQUENCE {
	return &TypeDescriptionSEQUENCE{tag: NewBerTag(0, 32, 16)}
}
