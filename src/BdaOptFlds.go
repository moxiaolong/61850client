package src

type BdaOptFlds struct {
	BdaBitString
}

func (f *BdaOptFlds) isSequenceNumber() bool {

}

func (f *BdaOptFlds) isReportTimestamp() bool {

}

func (f *BdaOptFlds) isDataSetName() bool {

}

func (f *BdaOptFlds) isBufferOverflow() bool {

}

func (f *BdaOptFlds) isEntryId() bool {

}

func (f *BdaOptFlds) isConfigRevision() bool {

}

func (f *BdaOptFlds) isSegmentation() bool {

}

func (f *BdaOptFlds) isDataReference() bool {

}

func (f *BdaOptFlds) isReasonForInclusion() bool {

}

func NewBdaOptFlds(*ObjectReference, interface{}) *BdaOptFlds {
	return &BdaOptFlds{}
}
