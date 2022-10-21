package src

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

func newMMSpdu() *MMSpdu {
	return &MMSpdu{}
}
func (s *MMSpdu) encode(stream *ReverseByteArrayOutputStream) {

}

func (s *MMSpdu) decode(is *ByteBufferInputStream) int {
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

func constructInitRequestPdu(int, int, int, int, []byte) *MMSpdu {
	return &MMSpdu{}
}
