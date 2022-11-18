package src

import (
	"bytes"
	"strconv"
)

type RejectPDU struct {
	originalInvokeID *Unsigned32
	tag              *BerTag
	rejectReason     *RejectReason
	code             []byte
}

func (p *RejectPDU) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0

	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += p.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		p.originalInvokeID = NewUnsigned32(0)
		vByteCount += p.originalInvokeID.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	p.rejectReason = NewRejectReason()
	numDecodedBytes = p.rejectReason.decode(is, berTag)
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

func (p *RejectPDU) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if p.code != nil {
		reverseOS.write(p.code)
		if withTag {
			return p.tag.encode(reverseOS) + len(p.code)
		}
		return len(p.code)
	}

	codeLength := 0
	codeLength += p.rejectReason.encode(reverseOS)

	if p.originalInvokeID != nil {
		codeLength += p.originalInvokeID.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.tag.encode(reverseOS)
	}

	return codeLength
}

func NewRejectPDU() *RejectPDU {
	return &RejectPDU{tag: NewBerTag(0, 32, 16)}

}
