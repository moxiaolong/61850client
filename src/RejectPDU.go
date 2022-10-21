package src

type RejectPDU struct {
}

func (p *RejectPDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewRejectPDU() *RejectPDU {
	return &RejectPDU{}

}
