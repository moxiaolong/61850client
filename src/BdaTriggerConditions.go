package src

type BdaTriggerConditions struct {
	BdaBitString
}

func NewBdaTriggerConditions(objectReference *ObjectReference, fc string) *BdaTriggerConditions {
	b := &BdaTriggerConditions{BdaBitString: *NewBdaBitString(objectReference, fc, "", 6, false, false)}

	b.basicType = TRIGGER_CONDITIONS
	b.setDefault()
	return b
}

func (s *BdaTriggerConditions) setDefault() {
	s.value = []byte{0x04}
}
