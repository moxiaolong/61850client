package src

import "strconv"

type BdaBoolean struct {
	BasicDataAttribute
	value  bool
	mirror *BdaBoolean
}

func (i *BdaBoolean) copy() ModelNodeI {
	boolean := NewBdaBoolean(i.ObjectReference, i.Fc, i.sAddr, i.dchg, i.dupd)
	boolean.value = i.value
	if i.mirror == nil {
		boolean.mirror = i
	} else {
		boolean.mirror = i.mirror
	}
	return boolean
}

func (i *BdaBoolean) setValueFromMmsDataObj(data *Data) {
	if data.bool == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: boolean")
	}
	i.value = data.bool.value
}

func (i *BdaBoolean) setDefault() {
	i.value = false
}
func (i *BdaBoolean) GetValueString() string {
	return strconv.FormatBool(i.value)
}
func NewBdaBoolean(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaBoolean {

	attribute := NewBasicDataAttribute(objectReference, fc, sAddr, dchg, dupd)
	attribute.basicType = BOOLEAN

	b := &BdaBoolean{BasicDataAttribute: *attribute}
	b.setDefault()
	return b
}
