package src

import "bytes"

type Data struct {
	visibleString *BerVisibleString
	bitString     *BerBitString
	Unsigned      *BerInteger
	bool          *BerBoolean
	octetString   *BerOctetString
	binaryTime    *TimeOfDay
	array         *Array
	integer       *BerInteger
	unsigned      *BerInteger
	floatingPoint *FloatingPoint
	mMSString     *MMSString
	utcTime       *UtcTime
	structure     *Structure
	code          []byte
}

func (d *Data) decode(is *bytes.Buffer, berTag *BerTag) int {

	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewEmptyBerTag()
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
		d.bool = NewBerBoolean()
		tlvByteCount += d.bool.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 4) {
		d.bitString = NewBerBitString(nil, nil, 0)
		tlvByteCount += d.bitString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 5) {
		d.integer = NewBerInteger(nil, 0)
		tlvByteCount += d.integer.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 6) {
		d.unsigned = NewBerInteger(nil, 0)
		tlvByteCount += d.unsigned.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 7) {
		d.floatingPoint = NewFloatingPoint()
		tlvByteCount += d.floatingPoint.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 9) {
		d.octetString = NewBerOctetString(nil)
		tlvByteCount += d.octetString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 10) {
		d.visibleString = NewBerVisibleString(nil)
		tlvByteCount += d.visibleString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 12) {
		d.binaryTime = NewTimeOfDay()
		tlvByteCount += d.binaryTime.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 16) {
		d.mMSString = NewMMSString()
		tlvByteCount += d.mMSString.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 17) {
		d.utcTime = NewUtcTime()
		tlvByteCount += d.utcTime.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (d *Data) encode(reverseOS *ReverseByteArrayOutputStream) int {
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
		// write tag: CONTEXT_CLASS, PRIMITIVE, 7
		reverseOS.writeByte(0x87)
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

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewData() *Data {
	return &Data{}
}
