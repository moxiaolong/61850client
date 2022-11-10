package src

import "bytes"

type TypeDescription struct {
	Structure     *Structure
	code          []byte
	array         *Array
	structure     *Structure
	bool          *BerNull
	bitString     *Integer32
	integer       *Unsigned8
	unsigned      *Unsigned8
	floatingPoint *FloatingPoint
	octetString   *Integer32
	visibleString *Integer32
	binaryTime    *BerBoolean
	mMSString     *Integer32
	utcTime       *BerNull
}

func (d *TypeDescription) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		d.array = NewArray()
		tlvByteCount += d.array.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 2) {
		d.structure = NewStructure()
		tlvByteCount += d.structure.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 3) {
		d.bool = NewBerNull()
		tlvByteCount += d.bool.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 4) {
		d.bitString = NewInteger32(0)
		tlvByteCount += d.bitString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 5) {
		d.integer = NewUnsigned8()
		tlvByteCount += d.integer.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 6) {
		d.unsigned = NewUnsigned8()
		tlvByteCount += d.unsigned.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 7) {
		d.floatingPoint = NewFloatingPoint()
		tlvByteCount += d.floatingPoint.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 9) {
		d.octetString = NewInteger32(0)
		tlvByteCount += d.octetString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 10) {
		d.visibleString = NewInteger32(0)
		tlvByteCount += d.visibleString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 12) {
		d.binaryTime = NewBerBoolean()
		tlvByteCount += d.binaryTime.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 16) {
		d.mMSString = NewInteger32(0)
		tlvByteCount += d.mMSString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 17) {
		d.utcTime = NewBerNull()
		tlvByteCount += d.utcTime.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (d *TypeDescription) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if d.code != nil {
		reverseOS.write(d.code)
		return len(d.code)
	}

	codeLength := 0
	if d.utcTime != nil {
		codeLength += d.utcTime.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 17
		reverseOS.writeByte(0x91)
		codeLength += 1
		return codeLength
	}

	if d.mMSString != nil {
		codeLength += d.mMSString.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 16
		reverseOS.writeByte(0x90)
		codeLength += 1
		return codeLength
	}

	if d.binaryTime != nil {
		codeLength += d.binaryTime.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 12
		reverseOS.writeByte(0x8C)
		codeLength += 1
		return codeLength
	}

	if d.visibleString != nil {
		codeLength += d.visibleString.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 10
		reverseOS.writeByte(0x8A)
		codeLength += 1
		return codeLength
	}

	if d.octetString != nil {
		codeLength += d.octetString.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.writeByte(0x89)
		codeLength += 1
		return codeLength
	}

	if d.floatingPoint != nil {
		codeLength += d.floatingPoint.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 7
		reverseOS.writeByte(0xA7)
		codeLength += 1
		return codeLength
	}

	if d.unsigned != nil {
		codeLength += d.unsigned.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 6
		reverseOS.writeByte(0x86)
		codeLength += 1
		return codeLength
	}

	if d.integer != nil {
		codeLength += d.integer.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 5
		reverseOS.writeByte(0x85)
		codeLength += 1
		return codeLength
	}

	if d.bitString != nil {
		codeLength += d.bitString.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 4
		reverseOS.writeByte(0x84)
		codeLength += 1
		return codeLength
	}

	if d.bool != nil {
		codeLength += d.bool.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
		return codeLength
	}

	if d.structure != nil {
		codeLength += d.structure.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
		return codeLength
	}

	if d.array != nil {
		codeLength += d.array.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func NewTypeDescription() *TypeDescription {
	return &TypeDescription{}
}
