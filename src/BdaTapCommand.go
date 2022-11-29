package src

type BdaTapCommand struct {
	BdaBitString
	mirror *BdaTapCommand
}

func (s *BdaTapCommand) copy() ModelNodeI {
	newCopy := NewBdaTapCommand(s.ObjectReference, s.Fc, s.sAddr, s.dchg, s.dupd)
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
func NewBdaTapCommand(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BdaTapCommand {
	bitString := NewBdaBitString(objectReference, fc, sAddr, 2, dchg, dupd)
	bitString.basicType = TAP_COMMAND
	b := &BdaTapCommand{BdaBitString: *bitString}
	b.setDefault()
	return b
}
func (b *BdaTapCommand) setDefault() {
	b.value = []byte{0x00}
}
