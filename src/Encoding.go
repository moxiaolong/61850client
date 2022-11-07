package src

import "bytes"

type Encoding struct {
	code           []byte
	tag            *BerTag
	singleASN1Type *BerAny
	octetAligned   *BerOctetString
	arbitrary      *BerBitString
}

func (e *Encoding) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 0) {

		length := NewBerLength()
		tlvByteCount += length.decode(is)
		e.singleASN1Type = NewBerAny(nil)
		tlvByteCount += e.singleASN1Type.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 1) {
		e.octetAligned = NewBerOctetString(nil)
		tlvByteCount += e.octetAligned.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 2) {
		e.arbitrary = NewBerBitString(nil, nil, 0)
		tlvByteCount += e.arbitrary.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}
func (e *Encoding) encode(reverseOS *ReverseByteArrayOutputStream, tag *BerTag) int {
	if e.code != nil {
		reverseOS.write(e.code)
		return len(e.code)
	}

	codeLength := 0
	sublength := 0

	if e.arbitrary != nil {
		codeLength += e.arbitrary.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
		return codeLength
	}

	if e.octetAligned != nil {
		codeLength += e.octetAligned.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	if e.singleASN1Type != nil {
		sublength = e.singleASN1Type.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func NewEncoding() *Encoding {
	return &Encoding{}
}
