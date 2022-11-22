package src

import "bytes"

type AlternateAccessSelection struct {
	selectAlternateAccess *SelectAlternateAccess
	selectAccess          *SelectAccess
	code                  []byte
}

func (s *AlternateAccessSelection) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if s.code != nil {
		reverseOS.write(s.code)
		return len(s.code)
	}

	codeLength := 0
	if s.selectAccess != nil {
		codeLength += s.selectAccess.encode(reverseOS)
		return codeLength
	}

	if s.selectAlternateAccess != nil {
		codeLength += s.selectAlternateAccess.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func (s *AlternateAccessSelection) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0

	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	numDecodedBytes := 0

	if berTag.equals(128, 32, 0) {
		s.selectAlternateAccess = NewSelectAlternateAccess()
		tlvByteCount += s.selectAlternateAccess.decode(is, false)
		return tlvByteCount
	}

	s.selectAccess = NewSelectAccess()
	numDecodedBytes = s.selectAccess.decode(is, berTag)
	if numDecodedBytes != 0 {
		return tlvByteCount + numDecodedBytes
	} else {
		s.selectAccess = nil
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func NewAlternateAccessSelection() *AlternateAccessSelection {
	return &AlternateAccessSelection{}
}
