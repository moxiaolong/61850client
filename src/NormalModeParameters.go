package src

import (
	"bytes"
	"strconv"
)

type CPAPPDUNormalModeParameters struct {
	tag                                     *BerTag
	protocolVersion                         *ProtocolVersion
	respondingPresentationSelector          *RespondingPresentationSelector
	presentationContextDefinitionResultList *PresentationContextDefinitionResultList
	presentationRequirements                *PresentationRequirements
	userSessionRequirements                 *UserSessionRequirements
	userData                                *UserData
	code                                    []byte
}

func (p *CPAPPDUNormalModeParameters) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += p.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val
	if lengthVal == 0 {
		return tlByteCount
	}
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		p.protocolVersion = NewProtocolVersion()
		vByteCount += p.protocolVersion.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 3) {
		p.respondingPresentationSelector = NewRespondingPresentationSelector(nil)
		vByteCount += p.respondingPresentationSelector.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 5) {
		p.presentationContextDefinitionResultList = NewPresentationContextDefinitionResultList()
		vByteCount += p.presentationContextDefinitionResultList.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 8) {
		p.presentationRequirements = NewPresentationRequirements()
		vByteCount += p.presentationRequirements.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 9) {
		p.userSessionRequirements = NewUserSessionRequirements()
		vByteCount += p.userSessionRequirements.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	p.userData = NewUserData()
	numDecodedBytes = p.userData.decode(is, berTag)
	if numDecodedBytes != 0 {
		vByteCount += numDecodedBytes
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		p.userData = nil
	}
	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw(
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (p *CPAPPDUNormalModeParameters) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if p.code != nil {
		reverseOS.write(p.code)
		if withTag {
			return p.tag.encode(reverseOS) + len(p.code)
		}
		return len(p.code)
	}

	codeLength := 0
	if p.userData != nil {
		codeLength += p.userData.encode(reverseOS)
	}

	if p.userSessionRequirements != nil {
		codeLength += p.userSessionRequirements.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.writeByte(0x89)
		codeLength += 1
	}

	if p.presentationRequirements != nil {
		codeLength += p.presentationRequirements.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 8
		reverseOS.writeByte(0x88)
		codeLength += 1
	}

	if p.presentationContextDefinitionResultList != nil {
		codeLength += p.presentationContextDefinitionResultList.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 5
		reverseOS.writeByte(0xA5)
		codeLength += 1
	}

	if p.respondingPresentationSelector != nil {
		codeLength += p.respondingPresentationSelector.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
	}

	if p.protocolVersion != nil {
		codeLength += p.protocolVersion.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += p.tag.encode(reverseOS)
	}

	return codeLength
}

func NewCPAPPDUNormalModeParameters() *CPAPPDUNormalModeParameters {
	return &CPAPPDUNormalModeParameters{tag: NewBerTag(0, 32, 16)}
}
