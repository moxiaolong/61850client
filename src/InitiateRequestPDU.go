package src

import "bytes"

type InitiateRequestPDU struct {
	LocalDetailCalling                *Integer32
	ProposedMaxServOutstandingCalling *Integer16
	ProposedMaxServOutstandingCalled  *Integer16
	ProposedDataStructureNestingLevel *Integer8
	InitRequestDetail                 *InitRequestDetail
}

func (p *InitiateRequestPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewInitiateRequestPDU() *InitiateRequestPDU {
	return &InitiateRequestPDU{}
}
