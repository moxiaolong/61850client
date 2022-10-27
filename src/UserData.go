package src

type UserData struct {
	FullyEncodedData *FullyEncodedData
}

func (t *UserData) encode(reverseOS *ReverseByteArrayOutputStream) int {

	codeLength := 0
	if t.FullyEncodedData != nil {
		codeLength += t.FullyEncodedData.encode(reverseOS, false)
		// write tag: APPLICATION_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0x61)
		codeLength += 1
		return codeLength
	}

	Throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func NewUserData() *UserData {
	return &UserData{}
}
