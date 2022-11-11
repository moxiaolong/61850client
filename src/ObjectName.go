package src

import "bytes"

type ObjectName struct {
	vmdSpecific    *Identifier
	domainSpecific *DomainSpecific
	aaSpecific     *Identifier
	code           []byte
}

func (n *ObjectName) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewEmptyBerTag()
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 0) {
		n.vmdSpecific = NewIdentifier()
		tlvByteCount += n.vmdSpecific.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 1) {
		n.domainSpecific = NewDomainSpecific()
		tlvByteCount += n.domainSpecific.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 2) {
		n.aaSpecific = NewIdentifier()
		tlvByteCount += n.aaSpecific.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (n *ObjectName) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if n.code != nil {
		reverseOS.write(n.code)
		return len(n.code)
	}

	codeLength := 0
	if n.aaSpecific != nil {
		codeLength += n.aaSpecific.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
		return codeLength
	}

	if n.domainSpecific != nil {
		codeLength += n.domainSpecific.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
		return codeLength
	}

	if n.vmdSpecific != nil {
		codeLength += n.vmdSpecific.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewObjectName() *ObjectName {
	return &ObjectName{}
}
