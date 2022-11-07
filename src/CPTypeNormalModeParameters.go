package src

import (
	"bytes"
	"strconv"
)

type CPTypeNormalModeParameters struct {
	CallingPresentationSelector       *CallingPresentationSelector
	CalledPresentationSelector        *CalledPresentationSelector
	PresentationContextDefinitionList *PresentationContextDefinitionList
	UserData                          *UserData
	tag                               *BerTag
	code                              []byte
	protocolVersion                   *ProtocolVersion
	callingPresentationSelector       *CallingPresentationSelector
	calledPresentationSelector        *CalledPresentationSelector
	presentationContextDefinitionList *PresentationContextDefinitionList
	defaultContextName                *DefaultContextName
	presentationRequirements          *PresentationRequirements
	userSessionRequirements           *UserSessionRequirements
	userData                          *UserData
}

func (t *CPTypeNormalModeParameters) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {

	if t.code != nil {
		reverseOS.write(t.code)
		if withTag {
			return t.tag.encode(reverseOS) + len(t.code)
		}
		return len(t.code)
	}

	codeLength := 0
	if t.userData != nil {
		codeLength += t.userData.encode(reverseOS)
	}

	if t.userSessionRequirements != nil {
		codeLength += t.userSessionRequirements.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.writeByte(0x89)
		codeLength += 1
	}

	if t.presentationRequirements != nil {
		codeLength += t.presentationRequirements.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 8
		reverseOS.writeByte(0x88)
		codeLength += 1
	}

	if t.defaultContextName != nil {
		codeLength += t.defaultContextName.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 6
		reverseOS.writeByte(0xA6)
		codeLength += 1
	}

	if t.presentationContextDefinitionList != nil {
		codeLength += t.presentationContextDefinitionList.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 4
		reverseOS.writeByte(0xA4)
		codeLength += 1
	}

	if t.calledPresentationSelector != nil {
		codeLength += t.calledPresentationSelector.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
	}

	if t.callingPresentationSelector != nil {
		codeLength += t.callingPresentationSelector.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
	}

	if t.protocolVersion != nil {
		codeLength += t.protocolVersion.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += t.tag.encode(reverseOS)
	}

	return codeLength
}

func (t *CPTypeNormalModeParameters) decode(is *bytes.Buffer, withTag bool) int {

	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0

	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += t.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	if lengthVal == 0 {
		return tlByteCount
	}
	vByteCount += berTag.decode(is)

	if berTag.equals(128, 0, 0) {
		t.protocolVersion = NewProtocolVersion()
		vByteCount += t.protocolVersion.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 1) {
		t.callingPresentationSelector = NewCallingPresentationSelector(nil)
		vByteCount += t.callingPresentationSelector.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 2) {
		t.calledPresentationSelector = NewCalledPresentationSelector(nil)
		vByteCount += t.calledPresentationSelector.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 4) {
		t.presentationContextDefinitionList = NewPresentationContextDefinitionList(nil)
		vByteCount += t.presentationContextDefinitionList.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 6) {
		t.defaultContextName = NewDefaultContextName()
		vByteCount += t.defaultContextName.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 8) {
		t.presentationRequirements = NewPresentationRequirements()
		vByteCount += t.presentationRequirements.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 9) {
		t.userSessionRequirements = NewUserSessionRequirements()
		vByteCount += t.userSessionRequirements.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	}

	t.userData = NewUserData()
	numDecodedBytes = t.userData.decode(is, berTag)
	if numDecodedBytes != 0 {
		vByteCount += numDecodedBytes
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		t.userData = nil
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

func NewCPTypeNormalModeParameters() *CPTypeNormalModeParameters {
	return &CPTypeNormalModeParameters{tag: NewBerTag(0, 32, 16)}
}
