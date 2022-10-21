package src

import "bytes"

type CPAPPDU struct {
}

func (c *CPAPPDU) decode(stream *bytes.Buffer) {

}

func NewCPAPPDU() *CPAPPDU {
	return &CPAPPDU{}
}
