package src

import "strconv"

type BdaInt8 struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt8
}

func (s *BdaInt8) getMmsDataObj() *Data {
	data := NewData()
	data.integer = NewBerInteger(nil, s.value)
	return data
}

func (s *BdaInt8) copy() ModelNodeI {
	newCopy := NewBdaInt8(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt8) setValueFromMmsDataObj(data *Data) {
	if data.integer == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: integer")
	}
	i.value = data.integer.value
}

func (i *BdaInt8) setDefault() {
	i.value = 0
}

func NewBdaInt8(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt8 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT8

	b := &BdaInt8{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}

func (i *BdaInt8) GetValueString() string {
	return strconv.Itoa(i.value)
}
