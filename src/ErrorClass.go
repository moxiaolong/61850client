package src

import (
	"bytes"
)

type ErrorClass struct {
	vmdState             *BerInteger
	applicationReference *BerInteger
	definition           *BerInteger
	resource             *BerInteger
	service              *BerInteger
	servicePreempt       *BerInteger
	timeResolution       *BerInteger
	access               *BerInteger
	initiate             *BerInteger
	conclude             *BerInteger
	cancel               *BerInteger
	file                 *BerInteger
	others               *BerInteger
	code                 []byte
}

func (c *ErrorClass) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 0) {
		c.vmdState = NewBerInteger(nil, 0)
		tlvByteCount += c.vmdState.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 1) {
		c.applicationReference = NewBerInteger(nil, 0)
		tlvByteCount += c.applicationReference.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 2) {
		c.definition = NewBerInteger(nil, 0)
		tlvByteCount += c.definition.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 3) {
		c.resource = NewBerInteger(nil, 0)
		tlvByteCount += c.resource.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 4) {
		c.service = NewBerInteger(nil, 0)
		tlvByteCount += c.service.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 5) {
		c.servicePreempt = NewBerInteger(nil, 0)
		tlvByteCount += c.servicePreempt.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 6) {
		c.timeResolution = NewBerInteger(nil, 0)
		tlvByteCount += c.timeResolution.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 7) {
		c.access = NewBerInteger(nil, 0)
		tlvByteCount += c.access.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 8) {
		c.initiate = NewBerInteger(nil, 0)
		tlvByteCount += c.initiate.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 9) {
		c.conclude = NewBerInteger(nil, 0)
		tlvByteCount += c.conclude.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 10) {
		c.cancel = NewBerInteger(nil, 0)
		tlvByteCount += c.cancel.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 11) {
		c.file = NewBerInteger(nil, 0)
		tlvByteCount += c.file.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 12) {
		c.others = NewBerInteger(nil, 0)
		tlvByteCount += c.others.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (c *ErrorClass) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if c.code != nil {
		reverseOS.write(c.code)
		return len(c.code)
	}

	codeLength := 0
	if c.others != nil {
		codeLength += c.others.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 12
		reverseOS.writeByte(0x8C)
		codeLength += 1
		return codeLength
	}

	if c.file != nil {
		codeLength += c.file.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 11
		reverseOS.writeByte(0x8B)
		codeLength += 1
		return codeLength
	}

	if c.cancel != nil {
		codeLength += c.cancel.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 10
		reverseOS.writeByte(0x8A)
		codeLength += 1
		return codeLength
	}

	if c.conclude != nil {
		codeLength += c.conclude.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.writeByte(0x89)
		codeLength += 1
		return codeLength
	}

	if c.initiate != nil {
		codeLength += c.initiate.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 8
		reverseOS.writeByte(0x88)
		codeLength += 1
		return codeLength
	}

	if c.access != nil {
		codeLength += c.access.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 7
		reverseOS.writeByte(0x87)
		codeLength += 1
		return codeLength
	}

	if c.timeResolution != nil {
		codeLength += c.timeResolution.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 6
		reverseOS.writeByte(0x86)
		codeLength += 1
		return codeLength
	}

	if c.servicePreempt != nil {
		codeLength += c.servicePreempt.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 5
		reverseOS.writeByte(0x85)
		codeLength += 1
		return codeLength
	}

	if c.service != nil {
		codeLength += c.service.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 4
		reverseOS.writeByte(0x84)
		codeLength += 1
		return codeLength
	}

	if c.resource != nil {
		codeLength += c.resource.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
		return codeLength
	}

	if c.definition != nil {
		codeLength += c.definition.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
		return codeLength
	}

	if c.applicationReference != nil {
		codeLength += c.applicationReference.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	if c.vmdState != nil {
		codeLength += c.vmdState.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewErrorClass() *ErrorClass {
	return &ErrorClass{}
}
