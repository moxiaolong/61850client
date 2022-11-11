package src

import "bytes"

type DomainName struct {
	code  []byte
	basic *BasicIdentifier
}

func (n *DomainName) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewEmptyBerTag()
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(0, 0, 26) {
		n.basic = NewBasicIdentifier()
		tlvByteCount += n.basic.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (n *DomainName) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if n.code != nil {
		reverseOS.write(n.code)
		return len(n.code)
	}

	codeLength := 0
	if n.basic != nil {
		codeLength += n.basic.encode(reverseOS, true)
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewDomainName() *DomainName {
	return &DomainName{}
}
