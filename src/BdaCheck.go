package src

type BdaCheck struct {
	BdaBitString
}

func NewBdaCheck(objectReference *ObjectReference) *BdaCheck {
	super := NewBdaBitString(objectReference, CO, "", 2, false, false)
	super.basicType = CHECK

	b := &BdaCheck{BdaBitString: *super}
	b.setDefault()
	return b
}
