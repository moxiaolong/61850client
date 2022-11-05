package src

import (
	"bytes"
)

type ErrorClass struct {
}

func (c *ErrorClass) decode(is *bytes.Buffer) int {

	int tlvByteCount = 0;
	boolean tagWasPassed = (berTag != nil);

	if (berTag == nil) {
		berTag = NewBerTag(0,0,0);
		tlvByteCount += berTag.decode(is);
	}

	if (berTag.equals(128, 0, 0)) {
		vmdState = NewBerInteger();
		tlvByteCount += vmdState.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 1)) {
		applicationReference = NewBerInteger();
		tlvByteCount += applicationReference.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 2)) {
		definition = NewBerInteger();
		tlvByteCount += definition.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 3)) {
		resource = NewBerInteger();
		tlvByteCount += resource.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 4)) {
		service = NewBerInteger();
		tlvByteCount += service.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 5)) {
		servicePreempt = NewBerInteger();
		tlvByteCount += servicePreempt.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 6)) {
		timeResolution = NewBerInteger();
		tlvByteCount += timeResolution.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 7)) {
		access = NewBerInteger();
		tlvByteCount += access.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 8)) {
		initiate = NewBerInteger();
		tlvByteCount += initiate.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 9)) {
		conclude = NewBerInteger();
		tlvByteCount += conclude.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 10)) {
		cancel = NewBerInteger();
		tlvByteCount += cancel.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 11)) {
		file = NewBerInteger();
		tlvByteCount += file.decode(is, false);
		return tlvByteCount;
	}

	if (berTag.equals(128, 0, 12)) {
		others = NewBerInteger();
		tlvByteCount += others.decode(is, false);
		return tlvByteCount;
	}

	if (tagWasPassed) {
		return 0;
	}

	throw("Error decoding CHOICE: Tag " + berTag + " matched to no item.");
}

func (c *ErrorClass) encode(os *ReverseByteArrayOutputStream) int {
	if (code != nil) {
		reverseOS.write(code);
		return code.length;
	}

	int codeLength = 0;
	if (others != nil) {
		codeLength += others.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 12
		reverseOS.write(0x8C);
		codeLength += 1;
		return codeLength;
	}

	if (file != nil) {
		codeLength += file.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 11
		reverseOS.write(0x8B);
		codeLength += 1;
		return codeLength;
	}

	if (cancel != nil) {
		codeLength += cancel.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 10
		reverseOS.write(0x8A);
		codeLength += 1;
		return codeLength;
	}

	if (conclude != nil) {
		codeLength += conclude.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.write(0x89);
		codeLength += 1;
		return codeLength;
	}

	if (initiate != nil) {
		codeLength += initiate.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 8
		reverseOS.write(0x88);
		codeLength += 1;
		return codeLength;
	}

	if (access != nil) {
		codeLength += access.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 7
		reverseOS.write(0x87);
		codeLength += 1;
		return codeLength;
	}

	if (timeResolution != nil) {
		codeLength += timeResolution.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 6
		reverseOS.write(0x86);
		codeLength += 1;
		return codeLength;
	}

	if (servicePreempt != nil) {
		codeLength += servicePreempt.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 5
		reverseOS.write(0x85);
		codeLength += 1;
		return codeLength;
	}

	if (service != nil) {
		codeLength += service.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 4
		reverseOS.write(0x84);
		codeLength += 1;
		return codeLength;
	}

	if (resource != nil) {
		codeLength += resource.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.write(0x83);
		codeLength += 1;
		return codeLength;
	}

	if (definition != nil) {
		codeLength += definition.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.write(0x82);
		codeLength += 1;
		return codeLength;
	}

	if (applicationReference != nil) {
		codeLength += applicationReference.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.write(0x81);
		codeLength += 1;
		return codeLength;
	}

	if (vmdState != nil) {
		codeLength += vmdState.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.write(0x80);
		codeLength += 1;
		return codeLength;
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.");
}

func NewErrorClass() *ErrorClass {
	return &ErrorClass{}
}
