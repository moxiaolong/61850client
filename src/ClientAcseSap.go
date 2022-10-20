package src

import (
	"bytes"
)

type ClientAcseSap struct {
	tSap *ClientTSap
}

func (s ClientAcseSap) associate(address string, port int, buffer *bytes.Buffer) *AcseAssociation {
	return nil
}

func newClientAcseSap() *ClientAcseSap {
	c := &ClientAcseSap{}
	c.tSap = NewClientTSap()
	return c
}
