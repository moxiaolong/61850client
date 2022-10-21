package src

type InitiateResponsePDU struct {
	LocalDetailCalled                   *LocalDetailCalled
	NegotiatedMaxServOutstandingCalling *NegotiatedMaxServOutstandingCalling
	NegotiatedMaxServOutstandingCalled  *NegotiatedMaxServOutstandingCalled
	NegotiatedDataStructureNestingLevel *NegotiatedDataStructureNestingLevel
	InitResponseDetail                  *InitResponseDetail
}

func (p *InitiateResponsePDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewInitiateResponsePDU() *InitiateResponsePDU {
	return &InitiateResponsePDU{}
}
