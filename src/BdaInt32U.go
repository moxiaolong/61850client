package src

import "strconv"

type BdaInt32U struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt32U
}

func (s *BdaInt32U) getMmsDataObj() *Data {
	data := NewData()
	data.integer = NewBerInteger(nil, s.value)
	return data
}

func (s *BdaInt32U) copy() ModelNodeI {
	newCopy := NewBdaInt32U(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt32U) setValueFromMmsDataObj(data *Data) {
	if data.unsigned == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: unsigned")
	}
	i.value = data.unsigned.value
}

func (i *BdaInt32U) setDefault() {
	i.value = 0
}
func NewBdaInt32U(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt32U {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT32U

	b := &BdaInt32U{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}

func (i *BdaInt32U) GetValueString() string {
	return strconv.Itoa(i.value)
}

func (s *BdaInt32U) setValue(atoi int) {
	s.value = atoi
}
