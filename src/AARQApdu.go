package src

import "math"

type AARQApdu struct {
	ApplicationContextName *BerObjectIdentifier
	CalledAPTitle          *APTitle
	CalledAEQualifier      *AEQualifier
	CallingAPTitle         *APTitle
	CallingAEQualifier     *AEQualifier
	UserInformation        *AssociationInformation
	Tag                    *BerTag
}

func (a *AARQApdu) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := 0
	sublength := 0

	if a.UserInformation != nil {
		codeLength += a.UserInformation.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 30
		reverseOS.writeByte(0xBE)
		codeLength += 1
	}

	if a.CallingAEQualifier != nil {
		sublength = a.CallingAEQualifier.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 7
		reverseOS.writeByte(0xA7)
		codeLength += 1
	}

	if a.CallingAPTitle != nil {
		sublength = a.CallingAPTitle.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 6
		reverseOS.writeByte(0xA6)
		codeLength += 1
	}

	if a.CalledAEQualifier != nil {
		sublength = a.CalledAEQualifier.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 3
		reverseOS.writeByte(0xA3)
		codeLength += 1
	}

	if a.CalledAPTitle != nil {
		sublength = a.CalledAPTitle.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
	}

	sublength = a.ApplicationContextName.encode(reverseOS, true)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += a.Tag.encode(reverseOS)
	}

	return codeLength
}

func NewAARQApdu() *AARQApdu {
	return &AARQApdu{Tag: NewBerTag(64, 32, 0)}
}

func encodeLength(reverseOS *ReverseByteArrayOutputStream, length int) int {
	if length <= 127 {
		reverseOS.writeByte(byte(length))
		return 1
	} else if length <= 255 {
		reverseOS.writeByte(byte(length))
		reverseOS.writeByte(129)
		return 2
	} else if length <= 65535 {
		reverseOS.writeByte(byte(length))
		reverseOS.writeByte(byte(length >> 8))
		reverseOS.writeByte(130)
		return 3
	} else if length <= 16777215 {
		reverseOS.writeByte(byte(length))
		reverseOS.writeByte(byte(length >> 8))
		reverseOS.writeByte(byte(length >> 16))
		reverseOS.writeByte(131)
		return 4
	} else {

		numLengthBytes := 1
		for (math.Pow(2.0, float64(8*numLengthBytes)) - 1.0) < float64(length) {
			numLengthBytes++
		}

		for i := 0; i < numLengthBytes; i++ {
			reverseOS.writeByte(byte(length >> 8 * i))

		}

		reverseOS.writeByte(byte(128 | numLengthBytes))
		return 1 + numLengthBytes
	}
}
