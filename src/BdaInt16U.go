package src

type BdaInt16U struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt16U
}

func (s *BdaInt16U) copy() ModelNodeI {
	newCopy := NewBdaInt16U(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt16U) setValueFromMmsDataObj(data *Data) {
	if data.unsigned == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: unsigned")
	}
	i.value = data.unsigned.value
}

func (i *BdaInt16U) setDefault() {
	i.value = 0
}
func NewBdaInt16U(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt16U {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT16U

	b := &BdaInt16U{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
