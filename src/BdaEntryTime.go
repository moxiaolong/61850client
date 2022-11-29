package src

import "strconv"

type BdaEntryTime struct {
	BasicDataAttribute
	value  []byte
	mirror *BdaEntryTime
}

func (f *BdaEntryTime) copy() ModelNodeI {
	newCopy := NewBdaEntryTime(f.ObjectReference, f.Fc, f.sAddr, f.dchg, f.dupd)
	valueCopy := make([]byte, 0)
	copy(valueCopy, f.value)
	newCopy.value = valueCopy
	if f.mirror == nil {
		newCopy.mirror = f
	} else {
		newCopy.mirror = f.mirror
	}
	return newCopy
}
func (t *BdaEntryTime) setValueFromMmsDataObj(data *Data) {
	if data.binaryTime == nil {
		throw("expected type: binary_time/EntryTime")
	}
	t.value = data.binaryTime.value
}

func (t *BdaEntryTime) setDefault() {
	t.value = make([]byte, 6)
}
func (t *BdaEntryTime) getStringValue() string {
	if len(t.value) != 6 {
		return strconv.Itoa(-1)
	}
	//TODO 需要测
	r := (((int(t.value[0]) & 0xff) << 24) + ((int(t.value[1]) & 0xff) << 16) + ((int(t.value[2]) & 0xff) << 8) + (int(t.value[3]) & 0xff) + (((int(t.value[4])&0xff)<<8)+(int(t.value[5])&0xff))*86400000) + 441763200000
	return strconv.Itoa(r)

}

func NewBdaEntryTime(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaEntryTime {
	basicDataAttribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	basicDataAttribute.basicType = ENTRY_TIME
	b := &BdaEntryTime{BasicDataAttribute: *basicDataAttribute}
	b.setDefault()
	return b
}
