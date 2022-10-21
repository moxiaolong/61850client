package src

import "bytes"

type RejectPDU struct {
}

func (p *RejectPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewRejectPDU() *RejectPDU {
	return &RejectPDU{}

}
