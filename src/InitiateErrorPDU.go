package src

import "bytes"

type InitiateErrorPDU struct {
	ErrorClass string
}

func (p *InitiateErrorPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewInitiateErrorPDU() *InitiateErrorPDU {
	return &InitiateErrorPDU{}

}
