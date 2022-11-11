package src

import (
	"bytes"
	"strconv"
)

type DefaultContextName struct {
	tag                *BerTag
	abstractSyntaxName *AbstractSyntaxName
	transferSyntaxName *TransferSyntaxName
	code               []byte
}

func (n *DefaultContextName) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0
	vByteCount := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += n.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		n.abstractSyntaxName = NewAbstractSyntaxName()
		vByteCount += n.abstractSyntaxName.decode(is, false)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 0, 1) {
		n.transferSyntaxName = NewTransferSyntaxName()
		vByteCount += n.transferSyntaxName.decode(is, false)
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

	throw(
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (n *DefaultContextName) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if n.code != nil {
		reverseOS.write(n.code)
		if withTag {
			return n.tag.encode(reverseOS) + len(n.code)
		}
		return len(n.code)
	}

	codeLength := 0
	codeLength += n.transferSyntaxName.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	codeLength += n.abstractSyntaxName.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += n.tag.encode(reverseOS)
	}

	return codeLength
}

func NewDefaultContextName() *DefaultContextName {
	return &DefaultContextName{tag: NewBerTag(0, 32, 16)}
}
