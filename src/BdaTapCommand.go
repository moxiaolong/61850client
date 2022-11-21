package src

type BdaTapCommand struct {
	BdaBitString
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
