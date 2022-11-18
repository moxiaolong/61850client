package src

import "bytes"

type PresentationDataValues struct {
	singleASN1Type *BerAny
	octetAligned   *BerOctetString
	arbitrary      *BerBitString
	code           []byte
}

func (v *PresentationDataValues) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if v.code != nil {
		reverseOS.write(v.code)
		return len(v.code)
	}
	codeLength := 0
	sublength := 0

	if v.arbitrary != nil {
		codeLength += v.encode(reverseOS)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
		return codeLength
	}

	if v.octetAligned != nil {
		codeLength += v.octetAligned.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	if v.singleASN1Type != nil {
		sublength = v.singleASN1Type.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return -1
}

func (v *PresentationDataValues) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewEmptyBerTag()
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 0) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		v.singleASN1Type = NewBerAny(nil)
		tlvByteCount += v.singleASN1Type.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 1) {
		v.octetAligned = NewBerOctetString(nil)
		tlvByteCount += v.octetAligned.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 2) {
		v.arbitrary = NewBerBitString(nil, nil, 0)
		tlvByteCount += v.arbitrary.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return -1
}

func NewPresentationDataValues() *PresentationDataValues {
	return &PresentationDataValues{}
}
