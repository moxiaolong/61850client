package src

import "bytes"

type TConnection struct {
}

func (c *TConnection) send(list [][]byte, offsets []int, lengths []int) {

}

func (c *TConnection) receive(buffer *bytes.Buffer) {

}

func NewTConnection() *TConnection {
	return &TConnection{}
}
