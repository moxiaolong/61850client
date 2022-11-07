package src

import (
	"bytes"
	"math"
	"strconv"
)

type AARQApdu struct {
	applicationContextName        *BerObjectIdentifier
	calledAPTitle                 *APTitle
	calledAEQualifier             *AEQualifier
	callingAPTitle                *APTitle
	callingAEQualifier            *AEQualifier
	userInformation               *AssociationInformation
	tag                           *BerTag
	code                          []byte
	protocolVersion               *BerBitString
	calledAPInvocationIdentifier  *APInvocationIdentifier
	calledAEInvocationIdentifier  *AEInvocationIdentifier
	callingAPInvocationIdentifier *APInvocationIdentifier
	callingAEInvocationIdentifier *AEInvocationIdentifier
	senderAcseRequirements        *ACSERequirements
	mechanismName                 *MechanismName
	callingAuthenticationValue    *AuthenticationValue
	applicationContextNameList    *ApplicationContextNameList
	implementationInformation     *ImplementationData
}

func (a *AARQApdu) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if a.code != nil {
		reverseOS.write(a.code)
		if withTag {
			return a.tag.encode(reverseOS) + len(a.code)
		}
		return len(a.code)
	}

	codeLength := 0
	sublength := 0

	if a.userInformation != nil {
		codeLength += a.userInformation.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 30
		reverseOS.writeByte(0xBE)
		codeLength += 1
	}

	if a.implementationInformation != nil {
		codeLength += a.implementationInformation.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 29
		reverseOS.writeByte(0x9D)
		codeLength += 1
	}

	if a.applicationContextNameList != nil {
		codeLength += a.applicationContextNameList.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 13
		reverseOS.writeByte(0xAD)
		codeLength += 1
	}

	if a.callingAuthenticationValue != nil {
		sublength = a.callingAuthenticationValue.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 12
		reverseOS.writeByte(0xAC)
		codeLength += 1
	}

	if a.mechanismName != nil {
		codeLength += a.mechanismName.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 11
		reverseOS.writeByte(0x8B)
		codeLength += 1
	}

	if a.senderAcseRequirements != nil {
		codeLength += a.senderAcseRequirements.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 10
		reverseOS.writeByte(0x8A)
		codeLength += 1
	}

	if a.callingAEInvocationIdentifier != nil {
		sublength = a.callingAEInvocationIdentifier.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 9
		reverseOS.writeByte(0xA9)
		codeLength += 1
	}

	if a.callingAPInvocationIdentifier != nil {
		sublength = a.callingAPInvocationIdentifier.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 8
		reverseOS.writeByte(0xA8)
		codeLength += 1
	}

	if a.callingAEQualifier != nil {
		sublength = a.callingAEQualifier.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 7
		reverseOS.writeByte(0xA7)
		codeLength += 1
	}

	if a.callingAPTitle != nil {
		sublength = a.callingAPTitle.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 6
		reverseOS.writeByte(0xA6)
		codeLength += 1
	}

	if a.calledAEInvocationIdentifier != nil {
		sublength = a.calledAEInvocationIdentifier.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 5
		reverseOS.writeByte(0xA5)
		codeLength += 1
	}

	if a.calledAPInvocationIdentifier != nil {
		sublength = a.calledAPInvocationIdentifier.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 4
		reverseOS.writeByte(0xA4)
		codeLength += 1
	}

	if a.calledAEQualifier != nil {
		sublength = a.calledAEQualifier.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 3
		reverseOS.writeByte(0xA3)
		codeLength += 1
	}

	if a.calledAPTitle != nil {
		sublength = a.calledAPTitle.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
	}

	sublength = a.applicationContextName.encode(reverseOS, true)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 1
	reverseOS.writeByte(0xA1)
	codeLength += 1

	if a.protocolVersion != nil {
		codeLength += a.protocolVersion.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += a.tag.encode(reverseOS)
	}

	return codeLength
}

func (a *AARQApdu) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0
	vByteCount := 0

	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += a.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		a.protocolVersion = NewBerBitString(nil, nil, 0)
		vByteCount += a.protocolVersion.decode(is, false)
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		vByteCount += length.decode(is)
		a.applicationContextName = NewBerObjectIdentifier(nil)
		vByteCount += a.applicationContextName.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 2) {
		vByteCount += length.decode(is)
		a.calledAPTitle = NewAPTitle()
		vByteCount += a.calledAPTitle.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 3) {
		vByteCount += length.decode(is)
		a.calledAEQualifier = NewAEQualifier()
		vByteCount += a.calledAEQualifier.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 4) {
		vByteCount += length.decode(is)
		a.calledAPInvocationIdentifier = NewAPInvocationIdentifier()
		vByteCount += a.calledAPInvocationIdentifier.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 5) {
		vByteCount += length.decode(is)
		a.calledAEInvocationIdentifier = NewAEInvocationIdentifier()
		vByteCount += a.calledAEInvocationIdentifier.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 6) {
		vByteCount += length.decode(is)
		a.callingAPTitle = NewAPTitle()
		vByteCount += a.callingAPTitle.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 7) {
		vByteCount += length.decode(is)
		a.callingAEQualifier = NewAEQualifier()
		vByteCount += a.callingAEQualifier.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 8) {
		vByteCount += length.decode(is)
		a.callingAPInvocationIdentifier = NewAPInvocationIdentifier()
		vByteCount += a.callingAPInvocationIdentifier.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 9) {
		vByteCount += length.decode(is)
		a.callingAEInvocationIdentifier = NewAEInvocationIdentifier()
		vByteCount += a.callingAEInvocationIdentifier.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 10) {
		a.senderAcseRequirements = NewACSERequirements()
		vByteCount += a.senderAcseRequirements.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 11) {
		a.mechanismName = NewMechanismName()
		vByteCount += a.mechanismName.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 12) {
		vByteCount += length.decode(is)
		a.callingAuthenticationValue = NewAuthenticationValue()
		vByteCount += a.callingAuthenticationValue.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 13) {
		a.applicationContextNameList = NewApplicationContextNameList()
		vByteCount += a.applicationContextNameList.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 29) {
		a.implementationInformation = NewImplementationData()
		vByteCount += a.implementationInformation.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 30) {
		a.userInformation = NewAssociationInformation()
		vByteCount += a.userInformation.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw("Unexpected end of sequence, length tag: ", strconv.Itoa(lengthVal), ", bytes decoded: ", strconv.Itoa(vByteCount))
	return 0
}

func NewAARQApdu() *AARQApdu {
	return &AARQApdu{tag: NewBerTag(64, 32, 0)}
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
