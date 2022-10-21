package src

type InitiateRequestPDU struct {
}

func (p *InitiateRequestPDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewInitiateRequestPDU() *InitiateRequestPDU {
	return &InitiateRequestPDU{}
}
