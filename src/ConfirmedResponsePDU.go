package src

import "bytes"

type ConfirmedResponsePDU struct {
	invokeID *Unsigned32
}

func (p *ConfirmedResponsePDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func (p *ConfirmedResponsePDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0
}

func NewConfirmedResponsePDU() *ConfirmedResponsePDU {
	return &ConfirmedResponsePDU{}

}
