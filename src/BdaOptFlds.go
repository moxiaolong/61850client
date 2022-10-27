package src

type BdaOptFlds struct {
	BdaBitString
}

func (f BdaOptFlds) isSequenceNumber() bool {

}

func (f BdaOptFlds) isReportTimestamp() bool {

}

func (f BdaOptFlds) isDataSetName() bool {

}

func NewBdaOptFlds(*ObjectReference, interface{}) *BdaOptFlds {
	return &BdaOptFlds{}
}
