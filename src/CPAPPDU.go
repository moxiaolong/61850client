package src

import "bytes"

type CPAPPDU struct {
}

func (c *CPAPPDU) decode(*bytes.Buffer) {

}

func NewCPAPPDU() *CPAPPDU {
	return &CPAPPDU{}
}
