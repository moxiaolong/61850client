package src

type BdaInt16 struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt16
}

func (s *BdaInt16) getMmsDataObj() *Data {
	data := NewData()
	data.integer = NewBerInteger(nil, s.value)
	return data
}

func (s *BdaInt16) copy() ModelNodeI {
	newCopy := NewBdaInt16(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt16) setValueFromMmsDataObj(data *Data) {
	if data.integer == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: integer")
	}
	i.value = data.integer.value
}

func (i *BdaInt16) setDefault() {
	i.value = 0
}

func (s *BdaInt16) setValue(atoi int) {
	s.value = atoi
}

func NewBdaInt16(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt16 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT16

	b := &BdaInt16{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
