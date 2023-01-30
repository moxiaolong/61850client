package src

import "strconv"

type BdaInt32 struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt32
}

func (s *BdaInt32) getMmsDataObj() *Data {
	data := NewData()
	data.integer = NewBerInteger(nil, s.value)
	return data
}

func (s *BdaInt32) copy() ModelNodeI {
	newCopy := NewBdaInt32(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt32) setValueFromMmsDataObj(data *Data) {
	if data.integer == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: integer")
	}
	i.value = data.integer.value
}

func (i *BdaInt32) setDefault() {
	i.value = 0
}

func NewBdaInt32(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt32 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT32

	b := &BdaInt32{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}

func (i *BdaInt32) GetValueString() string {
	return strconv.Itoa(i.value)
}

func (s *BdaInt32) setValue(atoi int) {
	s.value = atoi
}
