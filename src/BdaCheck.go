package src

type BdaCheck struct {
	BdaBitString
	mirror *BdaCheck
}

func (s *BdaCheck) copy() ModelNodeI {
	newCopy := NewBdaCheck(s.ObjectReference)
	valueCopy := make([]byte, 0)
	copy(valueCopy, s.value)
	newCopy.value = valueCopy
	if s.mirror == nil {
		newCopy.mirror = s
	} else {
		newCopy.mirror = s.mirror
	}
	return newCopy
}

func NewBdaCheck(objectReference *ObjectReference) *BdaCheck {
	super := NewBdaBitString(objectReference, CO, "", 2, false, false)
	super.basicType = CHECK

	b := &BdaCheck{BdaBitString: *super}
	b.setDefault()
	return b
}
