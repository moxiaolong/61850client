package src

import "bytes"

type AuthenticationValue struct {
	charstring *BerGraphicString
	bitstring  *BerBitString
	external   *Myexternal2
	code       []byte
}

func (v *AuthenticationValue) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 0) {
		v.charstring = NewBerGraphicString()
		tlvByteCount += v.charstring.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 1) {
		v.bitstring = NewBerBitString(nil, nil, 0)
		tlvByteCount += v.bitstring.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 2) {
		v.external = NewMyexternal2()
		tlvByteCount += v.external.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (v *AuthenticationValue) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if v.code != nil {
		reverseOS.write(v.code)
		return len(v.code)
	}

	codeLength := 0
	if v.external != nil {
		codeLength += v.external.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
		return codeLength
	}

	if v.bitstring != nil {
		codeLength += v.bitstring.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	if v.charstring != nil {
		codeLength += v.charstring.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewAuthenticationValue() *AuthenticationValue {
	return &AuthenticationValue{}
}
