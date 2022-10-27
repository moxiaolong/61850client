package src

type BdaEntryTime struct {
	BasicDataAttribute
}

func (t BdaEntryTime) setValueFromMmsDataObj(success *Data) {

}

func NewBdaEntryTime(*ObjectReference, interface{}, string, bool, bool) *BdaEntryTime {
	return &BdaEntryTime{}
}
