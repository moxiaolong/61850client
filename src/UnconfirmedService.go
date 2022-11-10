package src

import "bytes"

type UnconfirmedService struct {
	informationReport *InformationReport
	code              []byte
}

func (s *UnconfirmedService) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 0) {
		s.informationReport = NewInformationReport()
		tlvByteCount += s.informationReport.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (s *UnconfirmedService) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if s.code != nil {
		reverseOS.write(s.code)
		return len(s.code)
	}

	codeLength := 0
	if s.informationReport != nil {
		codeLength += s.informationReport.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}
	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func NewUnconfirmedService() *UnconfirmedService {
	return &UnconfirmedService{}
}
