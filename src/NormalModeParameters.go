package src

import (
	"bytes"
	"strconv"
)

type NormalModeParameters struct {
	tag                                     *BerTag
	protocolVersion                         *ProtocolVersion
	respondingPresentationSelector          *RespondingPresentationSelector
	presentationContextDefinitionResultList *PresentationContextDefinitionResultList
	presentationRequirements                *PresentationRequirements
	userSessionRequirements                 *UserSessionRequirements
	userData                                *UserData
}

func (p *NormalModeParameters) decode(is *bytes.Buffer, withTag bool) int {
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
		"Unexpected end of sequence, length tag: ", strconv.Itoa(lengthVal), ", bytes decoded: ", strconv.Itoa(vByteCount))
	return -1
}

func NewNormalModeParameters() *NormalModeParameters {
	return &NormalModeParameters{tag: NewBerTag(0, 32, 16)}
}
