package src

import "fmt"

type BdaFloat64 struct {
	BasicDataAttribute
	value  []byte
	mirror *BdaFloat64
}

func (s *BdaFloat64) getMmsDataObj() *Data {
	if s.value == nil {
		return nil
	}
	data := NewData()
	data.FloatingPoint = NewFloatingPoint(s.value)
	return data
}

func (s *BdaFloat64) copy() ModelNodeI {
	newCopy := NewBdaFloat64(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaFloat64) setValueFromMmsDataObj(data *Data) {

	if data.FloatingPoint == nil || len(data.FloatingPoint.value) != 9 {
		throw("ServiceError.TYPE_CONFLICT expected type: floating_point as an octet string of size 9")
	}
	i.value = data.FloatingPoint.value

}

func NewBdaFloat64(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaFloat64 {
	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = FLOAT64

	b := &BdaFloat64{BasicDataAttribute: *attribute}
	b.setDefault()
	return b

}

func (i *BdaFloat64) setDefault() {
	i.value = []byte{8, 0, 0, 0, 0}
}

func (i *BdaFloat64) GetValueString() string {
	//TODO 需要测
	return fmt.Sprintf("%s", i.value)
}
