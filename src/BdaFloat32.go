package src

import (
	"fmt"
	"unsafe"
)

type BdaFloat32 struct {
	BasicDataAttribute
	value  []byte
	mirror *BdaFloat32
}

func (s *BdaFloat32) copy() ModelNodeI {
	newCopy := NewBdaFloat32(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)

	newCopy.value = s.value
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func (i *BdaFloat32) GetValueString() string {
	//TODO 需要测
	t := ((0xff & int(i.value[1])) << 24) | ((0xff & int(i.value[2])) << 16) | ((0xff & int(i.value[3])) << 8) | ((0xff & int(i.value[4])) << 0)
	f := (*float32)(unsafe.Pointer(&t))
	return fmt.Sprintf("%f", float64(*f))
}
func (i *BdaFloat32) setValueFromMmsDataObj(data *Data) {
	if data.floatingPoint == nil || len(data.floatingPoint.value) != 5 {
		throw("ServiceError.TYPE_CONFLICT expected type: floating_point as an octet string of size 5")
	}
	i.value = data.floatingPoint.value
}

func (i *BdaFloat32) setDefault() {
	i.value = []byte{8, 0, 0, 0, 0}
}
func NewBdaFloat32(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaFloat32 {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = FLOAT32

	b := &BdaFloat32{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
