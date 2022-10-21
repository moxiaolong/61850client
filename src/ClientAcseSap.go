package src

import (
	"bytes"
)

type ClientAcseSap struct {
	tSap *ClientTSap
}

func (s *ClientAcseSap) associate(address string, port int, apdu *bytes.Buffer) *AcseAssociation {

	a := NewAcseAssociation(nil, []byte{0, 0, 0, 1})

	a.startAssociation(
		apdu,
		address,
		port,
		[]byte{0, 1},
		[]byte{0, 1},
		[]byte{0, 0, 0, 1},
		s.tSap,
		[]int{1, 1, 999, 1, 1},
		[]int{1, 1, 999, 1},
		12,
		12)

	defer func() {
		r := recover()
		if r != nil {
			a.disconnect()
			panic(r)
		}
	}()
	return a
}

func newClientAcseSap() *ClientAcseSap {
	c := &ClientAcseSap{}
	c.tSap = NewClientTSap()
	return c
}
