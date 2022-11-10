package src

import "bytes"

type VariableAccessSpecification struct {
	listOfVariable   *VariableDefs
	code             []byte
	variableListName *ObjectName
}

func (s *VariableAccessSpecification) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 0) {
		s.listOfVariable = NewVariableDefs()
		tlvByteCount += s.listOfVariable.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 1) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		s.variableListName = NewObjectName()
		tlvByteCount += s.variableListName.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (s *VariableAccessSpecification) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if s.code != nil {
		reverseOS.write(s.code)
		return len(s.code)
	}

	codeLength := 0
	sublength := 0

	if s.variableListName != nil {
		sublength = s.variableListName.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
		return codeLength
	}

	if s.listOfVariable != nil {
		codeLength += s.listOfVariable.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewVariableAccessSpecification() *VariableAccessSpecification {
	return &VariableAccessSpecification{}
}
