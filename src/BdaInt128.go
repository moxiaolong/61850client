package src

import (
	"strconv"
)

type BdaInt128 struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt128
}

func (s *BdaInt128) getMmsDataObj() *Data {
	data := NewData()
	data.integer = NewBerInteger(nil, s.value)
	return data
}

func (s *BdaInt128) copy() ModelNodeI {
	newCopy := NewBdaInt128(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt128) setValueFromMmsDataObj(data *Data) {
	if data.integer == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: integer")
	}
	i.value = data.integer.value
}

func (i *BdaInt128) setDefault() {
	i.value = 0
}

func NewBdaInt128(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt128 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT128

	b := &BdaInt128{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}

func (i *BdaInt128) GetValueString() string {
	return strconv.Itoa(i.value)
}

func (s *BdaInt128) setValue(atoi int) {

}
