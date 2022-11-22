package src

import "bytes"

type AccessSelection struct {
	code        []byte
	component   *AccessSelectionComponent
	index       *Unsigned32
	indexRange  *AccessSelectionIndexRange
	allElements *BerNull
}

func (s *AccessSelection) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if s.code != nil {
		reverseOS.write(s.code)
		return len(s.code)
	}

	codeLength := 0
	sublength := 0

	if s.allElements != nil {
		codeLength += s.allElements.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
		return codeLength
	}

	if s.indexRange != nil {
		codeLength += s.indexRange.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
		return codeLength
	}

	if s.index != nil {
		codeLength += s.index.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	if s.component != nil {
		sublength = s.component.encode(reverseOS)
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

func (s *AccessSelection) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 0) {

		length := NewBerLength()
		tlvByteCount += length.decode(is)
		s.component = NewAccessSelectionComponent()
		tlvByteCount += s.component.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 1) {
		s.index = NewUnsigned32(0)
		tlvByteCount += s.index.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 2) {
		s.indexRange = NewAccessSelectionIndexRange()
		tlvByteCount += s.indexRange.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 3) {
		s.allElements = NewBerNull()
		tlvByteCount += s.allElements.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func NewAccessSelection() *AccessSelection {
	return &AccessSelection{}
}
