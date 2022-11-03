package src

import (
	"bytes"
)

type RLRQApdu struct {
}

func (a RLRQApdu) decode(is *bytes.Buffer, b bool) int {

}

func NewRLRQApdu() *RLRQApdu {
	return &RLRQApdu{}
}
