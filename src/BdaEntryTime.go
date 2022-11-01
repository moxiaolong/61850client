package src

type BdaEntryTime struct {
	BasicDataAttribute
	value []byte
}

func (t *BdaEntryTime) setValueFromMmsDataObj(data *Data) {
	if data.binaryTime == nil {
		throw("expected type: binary_time/EntryTime")
	}
	t.value = data.binaryTime.value
}

func NewBdaEntryTime(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaEntryTime {
	basicDataAttribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	basicDataAttribute.basicType = ENTRY_TIME
	b := &BdaEntryTime{BasicDataAttribute: *basicDataAttribute}
	b.value = make([]byte, 6)
	return b
}
