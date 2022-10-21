package src

import "bytes"

type ConfirmedResponsePDU struct {
}

func (p *ConfirmedResponsePDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewConfirmedResponsePDU() *ConfirmedResponsePDU {
	return &ConfirmedResponsePDU{}

}
