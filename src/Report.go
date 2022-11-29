package src

type Report struct {
}

func NewReport(string, *int, *int, bool, string, *bool, *int, *BdaEntryTime, *BdaOctetString, []bool, []FcModelNodeI, []*BdaReasonForInclusion) *Report {
	return &Report{}
}
