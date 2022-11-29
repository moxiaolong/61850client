package src

type BdaInt64 struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt64
}

func (s *BdaInt64) copy() ModelNodeI {
	newCopy := NewBdaInt64(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt64) setValueFromMmsDataObj(data *Data) {
	if data.integer == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: integer")
	}
	i.value = data.integer.value
}

func (i *BdaInt64) setDefault() {
	i.value = 0
}

func NewBdaInt64(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt64 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT64

	b := &BdaInt64{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
