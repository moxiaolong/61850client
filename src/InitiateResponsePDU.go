package src

import "bytes"

type InitiateResponsePDU struct {
	LocalDetailCalled                   *LocalDetailCalled
	NegotiatedMaxServOutstandingCalling *NegotiatedMaxServOutstandingCalling
	NegotiatedMaxServOutstandingCalled  *NegotiatedMaxServOutstandingCalled
	NegotiatedDataStructureNestingLevel *NegotiatedDataStructureNestingLevel
	InitResponseDetail                  *InitResponseDetail
}

func (p *InitiateResponsePDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func (p *InitiateResponsePDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0
}

func NewInitiateResponsePDU() *InitiateResponsePDU {
	return &InitiateResponsePDU{}
}
