package src

import "bytes"

type UnconfirmedPDU struct {
	Service *UnconfirmedService
}

func (p *UnconfirmedPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func (p *UnconfirmedPDU) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0
}

func NewUnconfirmedPDU() *UnconfirmedPDU {
	return &UnconfirmedPDU{}
}
