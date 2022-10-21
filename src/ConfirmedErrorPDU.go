package src

import "bytes"

type ConfirmedErrorPDU struct {
}

func (p *ConfirmedErrorPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewConfirmedErrorPDU() *ConfirmedErrorPDU {
	return &ConfirmedErrorPDU{}
}
