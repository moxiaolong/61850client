package src

type ConfirmedResponsePDU struct {
}

func (p *ConfirmedResponsePDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewConfirmedResponsePDU() *ConfirmedResponsePDU {
	return &ConfirmedResponsePDU{}

}
