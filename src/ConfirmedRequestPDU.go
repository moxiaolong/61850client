package src

type ConfirmedRequestPDU struct {
}

func (p *ConfirmedRequestPDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewConfirmedRequestPDU() *ConfirmedRequestPDU {
	return &ConfirmedRequestPDU{}
}
