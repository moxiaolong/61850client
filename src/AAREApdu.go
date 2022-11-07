package src

import (
	"bytes"
	"strconv"
)

type AAREApdu struct {
	UserInformation                  *UserInformation
	code                             []byte
	tag                              *BerTag
	protocolVersion                  *BerBitString
	applicationContextName           *BerObjectIdentifier
	result                           *AssociateResult
	resultSourceDiagnostic           *AssociateSourceDiagnostic
	respondingAPTitle                *APTitle
	respondingAEQualifier            *AEQualifier
	respondingAPInvocationIdentifier *APInvocationIdentifier
	respondingAEInvocationIdentifier *AEInvocationIdentifier
	responderAcseRequirements        *ACSERequirements
	mechanismName                    *MechanismName
	respondingAuthenticationValue    *AuthenticationValue
	applicationContextNameList       *ApplicationContextNameList
	implementationInformation        *ImplementationData
	userInformation                  *AssociationInformation
}

func (a *AAREApdu) decode(is *bytes.Buffer, withTag bool) int {
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
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 2) {
		vByteCount += length.decode(is)
		a.result = NewAssociateResult()
		vByteCount += a.result.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 3) {
		vByteCount += length.decode(is)
		a.resultSourceDiagnostic = NewAssociateSourceDiagnostic()
		vByteCount += a.resultSourceDiagnostic.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if berTag.equals(128, 32, 4) {
		vByteCount += length.decode(is)
		a.respondingAPTitle = NewAPTitle()
		vByteCount += a.respondingAPTitle.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 5) {
		vByteCount += length.decode(is)
		a.respondingAEQualifier = NewAEQualifier()
		vByteCount += a.respondingAEQualifier.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 6) {
		vByteCount += length.decode(is)
		a.respondingAPInvocationIdentifier = NewAPInvocationIdentifier()
		vByteCount += a.respondingAPInvocationIdentifier.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 7) {
		vByteCount += length.decode(is)
		a.respondingAEInvocationIdentifier = NewAEInvocationIdentifier()
		vByteCount += a.respondingAEInvocationIdentifier.decode(is, true)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 8) {
		a.responderAcseRequirements = NewACSERequirements()
		vByteCount += a.responderAcseRequirements.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 9) {
		a.mechanismName = NewMechanismName()
		vByteCount += a.mechanismName.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 10) {
		vByteCount += length.decode(is)
		a.respondingAuthenticationValue = NewAuthenticationValue()
		vByteCount += a.respondingAuthenticationValue.decode(is, nil)
		vByteCount += length.readEocIfIndefinite(is)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 11) {
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

func (a *AAREApdu) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
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
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 11
		reverseOS.writeByte(0xAB)
		codeLength += 1
	}

	if a.respondingAuthenticationValue != nil {
		sublength = a.respondingAuthenticationValue.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 10
		reverseOS.writeByte(0xAA)
		codeLength += 1
	}

	if a.mechanismName != nil {
		codeLength += a.mechanismName.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.writeByte(0x89)
		codeLength += 1
	}

	if a.responderAcseRequirements != nil {
		codeLength += a.responderAcseRequirements.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 8
		reverseOS.writeByte(0x88)
		codeLength += 1
	}

	if a.respondingAEInvocationIdentifier != nil {
		sublength = a.respondingAEInvocationIdentifier.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 7
		reverseOS.writeByte(0xA7)
		codeLength += 1
	}

	if a.respondingAPInvocationIdentifier != nil {
		sublength = a.respondingAPInvocationIdentifier.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 6
		reverseOS.writeByte(0xA6)
		codeLength += 1
	}

	if a.respondingAEQualifier != nil {
		sublength = a.respondingAEQualifier.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 5
		reverseOS.writeByte(0xA5)
		codeLength += 1
	}

	if a.respondingAPTitle != nil {
		sublength = a.respondingAPTitle.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 4
		reverseOS.writeByte(0xA4)
		codeLength += 1
	}

	sublength = a.resultSourceDiagnostic.encode(reverseOS)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 3
	reverseOS.writeByte(0xA3)
	codeLength += 1

	sublength = a.result.encode(reverseOS, true)
	codeLength += sublength
	codeLength += encodeLength(reverseOS, sublength)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 2
	reverseOS.writeByte(0xA2)
	codeLength += 1

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

func NewAAREApdu() *AAREApdu {
	return &AAREApdu{tag: NewBerTag(64, 32, 1)}
}
