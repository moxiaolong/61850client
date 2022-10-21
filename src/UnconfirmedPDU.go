package src

type UnconfirmedPDU struct {
}

func (p *UnconfirmedPDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewUnconfirmedPDU() *UnconfirmedPDU {
	return &UnconfirmedPDU{}
}
