package src

import "bytes"

type ConfirmedErrorPDU struct {
	invokeID *Unsigned32
}

func (p *ConfirmedErrorPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func (p *ConfirmedErrorPDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0
}

func NewConfirmedErrorPDU() *ConfirmedErrorPDU {
	return &ConfirmedErrorPDU{}
}
