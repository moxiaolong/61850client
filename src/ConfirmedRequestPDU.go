package src

import "bytes"

type ConfirmedRequestPDU struct {
}

func (p *ConfirmedRequestPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewConfirmedRequestPDU() *ConfirmedRequestPDU {
	return &ConfirmedRequestPDU{}
}
