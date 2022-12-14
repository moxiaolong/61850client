package src

import "bytes"

type UserData struct {
	simplyEncodedData *SimplyEncodedData
	fullyEncodedData  *FullyEncodedData
	code              []byte
}

func (t *UserData) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if t.code != nil {
		reverseOS.write(t.code)
		return len(t.code)
	}
	codeLength := 0
	if t.fullyEncodedData != nil {
		codeLength += t.fullyEncodedData.encode(reverseOS, false)
		// writeByte tag: APPLICATION_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0x61)
		codeLength += 1
		return codeLength
	}
	if t.simplyEncodedData != nil {
		codeLength += t.simplyEncodedData.encode(reverseOS, false)
		// write tag: APPLICATION_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x40)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return -1
}

func (t *UserData) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewEmptyBerTag()
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

	throw("Error decoding WriteResponseCHOICE: tag ", berTag.toString(), " matched to no item.")
	return 0
}

func NewUserData() *UserData {
	return &UserData{}
}
