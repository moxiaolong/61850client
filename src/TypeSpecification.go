package src

import "bytes"

type TypeSpecification struct {
	typeDescription *TypeDescription
	code            []byte
}

func (s *TypeSpecification) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	numDecodedBytes := 0
	s.typeDescription = NewTypeDescription()
	numDecodedBytes = s.typeDescription.decode(is, berTag)
	if numDecodedBytes != 0 {
		return tlvByteCount + numDecodedBytes
	} else {
		s.typeDescription = nil
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (s *TypeSpecification) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if s.code != nil {
		reverseOS.write(s.code)
		return len(s.code)
	}

	codeLength := 0
	if s.typeDescription != nil {
		codeLength += s.typeDescription.encode(reverseOS)
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewTypeSpecification() *TypeSpecification {
	return &TypeSpecification{}
}
