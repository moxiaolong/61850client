package src

import "bytes"

type VariableSpecification struct {
	name *ObjectName
	code []byte
}

func (s *VariableSpecification) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if s.code != nil {
		reverseOS.write(s.code)
		return len(s.code)
	}

	codeLength := 0
	sublength := 0

	if s.name != nil {
		sublength = s.name.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func (s *VariableSpecification) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 0) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		s.name = NewObjectName()
		tlvByteCount += s.name.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func NewVariableSpecification() *VariableSpecification {
	return &VariableSpecification{}
}
