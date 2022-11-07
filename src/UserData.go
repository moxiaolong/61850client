package src

import "bytes"

type UserData struct {
	simplyEncodedData *SimplyEncodedData
	fullyEncodedData  *FullyEncodedData
}

func (t *UserData) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if code != nil {
		reverseOS.writeByte(code)
		return code.length
	}
	codeLength := 0
	if t.fullyEncodedData != nil {
		codeLength += t.fullyEncodedData.encode(reverseOS, false)
		// writeByte tag: APPLICATION_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0x61)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func (t *UserData) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(64, 0, 0) {
		t.simplyEncodedData = NewSimplyEncodedData()
		tlvByteCount += t.simplyEncodedData.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(64, 32, 1) {
		t.fullyEncodedData = NewFullyEncodedData()
		tlvByteCount += t.fullyEncodedData.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: tag ", berTag.toString(), " matched to no item.")
	return 0
}

func NewUserData() *UserData {
	return &UserData{}
}
