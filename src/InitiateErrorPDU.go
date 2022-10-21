package src

type InitiateErrorPDU struct {
	ErrorClass string
}

func (p *InitiateErrorPDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewInitiateErrorPDU() *InitiateErrorPDU {
	return &InitiateErrorPDU{}

}
