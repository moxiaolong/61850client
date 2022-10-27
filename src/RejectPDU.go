package src

import "bytes"

type RejectPDU struct {
	OriginalInvokeID *Unsigned32
}

func (p *RejectPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func (p *RejectPDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0
}

func NewRejectPDU() *RejectPDU {
	return &RejectPDU{}

}
