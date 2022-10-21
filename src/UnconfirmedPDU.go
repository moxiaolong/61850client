package src

import "bytes"

type UnconfirmedPDU struct {
}

func (p *UnconfirmedPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewUnconfirmedPDU() *UnconfirmedPDU {
	return &UnconfirmedPDU{}
}
