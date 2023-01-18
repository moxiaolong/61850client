package src

type BdaTriggerConditions struct {
	BdaBitString
	mirror *BdaTriggerConditions
}

func (s *BdaTriggerConditions) copy() ModelNodeI {
	newCopy := NewBdaTriggerConditions(s.ObjectReference, s.Fc)
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

func NewBdaTriggerConditions(objectReference *ObjectReference, fc string) *BdaTriggerConditions {
	b := &BdaTriggerConditions{BdaBitString: *NewBdaBitString(objectReference, fc, "", 6, false, false)}

	b.basicType = TRIGGER_CONDITIONS
	b.setDefault()
	return b
}

func (s *BdaTriggerConditions) setDefault() {
	s.value = []byte{0x04}
}

func (s *BdaTriggerConditions) setDataChange(b bool) {
	//TODO
}

func (s *BdaTriggerConditions) setQualityChange(b bool) {

}

func (s *BdaTriggerConditions) setDataUpdate(b bool) {

}

func (s *BdaTriggerConditions) setIntegrity(b bool) {

}

func (s *BdaTriggerConditions) setGeneralInterrogation(b bool) {

}
