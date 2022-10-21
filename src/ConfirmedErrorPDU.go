package src

type ConfirmedErrorPDU struct {
}

func (p *ConfirmedErrorPDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewConfirmedErrorPDU() *ConfirmedErrorPDU {
	return &ConfirmedErrorPDU{}
}
