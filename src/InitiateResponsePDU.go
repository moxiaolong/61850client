package src

type InitiateResponsePDU struct {
	LocalDetailCalled                   *LocalDetailCalled
	NegotiatedMaxServOutstandingCalling *NegotiatedMaxServOutstandingCalling
	NegotiatedMaxServOutstandingCalled  *NegotiatedMaxServOutstandingCalled
	NegotiatedDataStructureNestingLevel *NegotiatedDataStructureNestingLevel
	InitResponseDetail                  *InitResponseDetail
}
