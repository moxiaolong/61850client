package src

import "bytes"

type ObjectScope struct {
	vmdSpecific    *BerNull
	domainSpecific *Identifier
	aaSpecific     *BerNull
	code           []byte
}

func (s *ObjectScope) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 0) {
		s.vmdSpecific = NewBerNull()
		tlvByteCount += s.vmdSpecific.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 1) {
		s.domainSpecific = NewIdentifier()
		tlvByteCount += s.domainSpecific.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 2) {
		s.aaSpecific = NewBerNull()
		tlvByteCount += s.aaSpecific.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (s *ObjectScope) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if s.code != nil {
		reverseOS.write(s.code)
		return len(s.code)
	}

	codeLength := 0
	if s.aaSpecific != nil {
		codeLength += s.aaSpecific.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
		return codeLength
	}

	if s.domainSpecific != nil {
		codeLength += s.domainSpecific.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	if s.vmdSpecific != nil {
		codeLength += s.vmdSpecific.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewObjectScope() *ObjectScope {
	return &ObjectScope{}
}
