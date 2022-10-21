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
	InitiateRequestPDU   *InitiateRequestPDU
}

func NewMMSpdu() *MMSpdu {
	return &MMSpdu{}
}
func (s *MMSpdu) encode(stream *ReverseByteArrayOutputStream) {

}

func (s *MMSpdu) decode(is *bytes.Buffer) int {
	tlvByteCount := 0

	berTag := NewBerTag()
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
	initRequestDetail.ProposedVersionNumber = NewInteger16([]byte{0x01, 0x01})
	initRequestDetail.ProposedParameterCBB = NewParameterSupportOptions([]byte{0x03, 0x05, 0xf1, 0x00})
	initRequestDetail.ServicesSupportedCalling = NewServiceSupportOptions(servicesSupportedCalling, 85)

	initiateRequestPdu := NewInitiateRequestPDU()
	initiateRequestPdu.LocalDetailCalling = NewInteger32(proposedMaxPduSize)
	initiateRequestPdu.ProposedMaxServOutstandingCalling = NewInteger16Int(proposedMaxServOutstandingCalling)

	initiateRequestPdu.ProposedMaxServOutstandingCalled = NewInteger16Int(proposedMaxServOutstandingCalled)
	initiateRequestPdu.ProposedDataStructureNestingLevel = NewInteger8(proposedDataStructureNestingLevel)
	initiateRequestPdu.InitRequestDetail = initRequestDetail

	initiateRequestMMSpdu := NewMMSpdu()
	initiateRequestMMSpdu.InitiateRequestPDU = initiateRequestPdu

	return initiateRequestMMSpdu
}
