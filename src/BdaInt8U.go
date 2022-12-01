package src

type BdaInt8U struct {
	BasicDataAttribute
	value  int
	mirror *BdaInt8U
}

func (s *BdaInt8U) getMmsDataObj() *Data {
	data := NewData()
	data.integer = NewBerInteger(nil, s.value)
	return data
}

func (s *BdaInt8U) copy() ModelNodeI {
	newCopy := NewBdaInt8U(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaInt8U) setValueFromMmsDataObj(data *Data) {
	if data.unsigned == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: unsigned")
	}
	i.value = data.unsigned.value
}

func (i *BdaInt8U) setDefault() {
	i.value = 0
}
func NewBdaInt8U(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaInt8U {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = INT8U

	b := &BdaInt8U{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
