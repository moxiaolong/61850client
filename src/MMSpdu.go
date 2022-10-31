package src

import "bytes"

type MMSpdu struct {
	confirmedRequestPDU  *ConfirmedRequestPDU
	confirmedResponsePDU *ConfirmedResponsePDU
	confirmedErrorPDU    *ConfirmedErrorPDU
	unconfirmedPDU       *UnconfirmedPDU
	rejectPDU            *RejectPDU
	initiateRequestPDU   *InitiateRequestPDU
	initiateResponsePDU  *InitiateResponsePDU
	initiateErrorPDU     *InitiateErrorPDU
	concludeRequestPDU   *ConcludeRequestPDU
}

func NewMMSpdu() *MMSpdu {
	return &MMSpdu{}
}
func (s *MMSpdu) encode(reverseOS *ReverseByteArrayOutputStream) int {

	codeLength := 0
	if s.concludeRequestPDU != nil {
		codeLength += s.concludeRequestPDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 11
		reverseOS.writeByte(0x8B)
		codeLength += 1
		return codeLength
	}

	if s.initiateErrorPDU != nil {
		codeLength += s.initiateErrorPDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 10
		reverseOS.writeByte(0xAA)
		codeLength += 1
		return codeLength
	}

	if s.initiateResponsePDU != nil {
		codeLength += s.initiateResponsePDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 9
		reverseOS.writeByte(0xA9)
		codeLength += 1
		return codeLength
	}

	if s.initiateRequestPDU != nil {
		codeLength += s.initiateRequestPDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 8
		reverseOS.writeByte(0xA8)
		codeLength += 1
		return codeLength
	}

	if s.rejectPDU != nil {
		codeLength += s.rejectPDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 4
		reverseOS.writeByte(0xA4)
		codeLength += 1
		return codeLength
	}

	if s.unconfirmedPDU != nil {
		codeLength += s.unconfirmedPDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 3
		reverseOS.writeByte(0xA3)
		codeLength += 1
		return codeLength
	}

	if s.confirmedErrorPDU != nil {
		codeLength += s.confirmedErrorPDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
		return codeLength
	}

	if s.confirmedResponsePDU != nil {
		codeLength += s.confirmedResponsePDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
		return codeLength
	}

	if s.confirmedRequestPDU != nil {
		codeLength += s.confirmedRequestPDU.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func (s *MMSpdu) decode(is *bytes.Buffer) int {
	tlvByteCount := 0

	berTag := NewBerTag(0, 0, 0)
	tlvByteCount += berTag.decode(is)

	if berTag.equals(128, 32, 0) {
		s.confirmedRequestPDU = NewConfirmedRequestPDU()
		tlvByteCount += s.confirmedRequestPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 1) {
		s.confirmedResponsePDU = NewConfirmedResponsePDU()
		tlvByteCount += s.confirmedResponsePDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 2) {
		s.confirmedErrorPDU = NewConfirmedErrorPDU()
		tlvByteCount += s.confirmedErrorPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 3) {
		s.unconfirmedPDU = NewUnconfirmedPDU()
		tlvByteCount += s.unconfirmedPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 4) {
		s.rejectPDU = NewRejectPDU()
		tlvByteCount += s.rejectPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 8) {
		s.initiateRequestPDU = NewInitiateRequestPDU()
		tlvByteCount += s.initiateRequestPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 9) {
		s.initiateResponsePDU = NewInitiateResponsePDU()
		tlvByteCount += s.initiateResponsePDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 10) {
		s.initiateErrorPDU = NewInitiateErrorPDU()
		tlvByteCount += s.initiateErrorPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 11) {
		s.concludeRequestPDU = NewConcludeRequestPDU()
		tlvByteCount += s.concludeRequestPDU.decode(is, false)
		return tlvByteCount
	}

	return 0

}

func constructInitRequestPdu(proposedMaxPduSize int, proposedMaxServOutstandingCalling int, proposedMaxServOutstandingCalled int, proposedDataStructureNestingLevel int, servicesSupportedCalling []byte) *MMSpdu {
	initRequestDetail := NewInitRequestDetail()
	initRequestDetail.proposedVersionNumber = NewInteger16([]byte{0x01, 0x01}, 0)
	initRequestDetail.proposedParameterCBB = NewParameterSupportOptions([]byte{0x03, 0x05, 0xf1, 0x00})
	initRequestDetail.servicesSupportedCalling = NewServiceSupportOptions(servicesSupportedCalling, 85)

	initiateRequestPdu := NewInitiateRequestPDU()
	initiateRequestPdu.LocalDetailCalling = NewInteger32(proposedMaxPduSize)
	initiateRequestPdu.ProposedMaxServOutstandingCalling = NewInteger16(nil, proposedMaxServOutstandingCalling)

	initiateRequestPdu.ProposedMaxServOutstandingCalled = NewInteger16(nil, proposedMaxServOutstandingCalled)
	initiateRequestPdu.ProposedDataStructureNestingLevel = NewInteger8(proposedDataStructureNestingLevel)
	initiateRequestPdu.InitRequestDetail = initRequestDetail

	initiateRequestMMSpdu := NewMMSpdu()
	initiateRequestMMSpdu.initiateRequestPDU = initiateRequestPdu

	return initiateRequestMMSpdu
}
