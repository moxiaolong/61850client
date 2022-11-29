package src

import "strconv"

type BdaTimestamp struct {
	BasicDataAttribute
	value  []byte
	mirror *BdaTimestamp
}

func (f *BdaTimestamp) copy() ModelNodeI {
	newCopy := NewBdaTimestamp(f.ObjectReference, f.Fc, f.sAddr, f.dchg, f.dupd)
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

func (t *BdaTimestamp) setValueFromMmsDataObj(data *Data) {
	if data.utcTime == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: utc_time/timestamp")
	}
	t.value = data.utcTime.value
}

func (t *BdaTimestamp) setDefault() {
	t.value = make([]byte, 8)
}

func NewBdaTimestamp(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaTimestamp {

	b := &BdaTimestamp{BasicDataAttribute: *NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)}
	b.basicType = TIMESTAMP
	b.setDefault()
	return b
}

func (t *BdaTimestamp) GetValueString() string {
	r := (0xff&int(t.value[0]))<<24 | (0xff&int(t.value[1]))<<16 | (0xff&int(t.value[2]))<<8 | (0xff & int(t.value[3]))

	return strconv.Itoa(r)
}
