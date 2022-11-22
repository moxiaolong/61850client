package src

import "bytes"

type SelectAccess struct {
	code        []byte
	component   *SelectAccessComponent
	index       *Unsigned32
	indexRange  *SelectAccessIndexRange
	allElements *BerNull
}

func (a *SelectAccess) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if a.code != nil {
		reverseOS.write(a.code)
		return len(a.code)
	}

	codeLength := 0
	sublength := 0

	if a.allElements != nil {
		codeLength += a.allElements.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 4
		reverseOS.writeByte(0x84)
		codeLength += 1
		return codeLength
	}

	if a.indexRange != nil {
		codeLength += a.indexRange.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 3
		reverseOS.writeByte(0xA3)
		codeLength += 1
		return codeLength
	}

	if a.index != nil {
		codeLength += a.index.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
		return codeLength
	}

	if a.component != nil {
		sublength = a.component.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func (a *SelectAccess) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0

	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {

		length := NewBerLength()
		tlvByteCount += length.decode(is)
		a.component = NewSelectAccessComponent()
		tlvByteCount += a.component.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 2) {
		a.index = NewUnsigned32(0)
		tlvByteCount += a.index.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 3) {
		a.indexRange = NewSelectAccessIndexRange()
		tlvByteCount += a.indexRange.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 4) {
		a.allElements = NewBerNull()
		tlvByteCount += a.allElements.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func NewSelectAccess() *SelectAccess {
	return &SelectAccess{}
}
