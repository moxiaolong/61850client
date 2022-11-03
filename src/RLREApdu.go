package src

import (
	"bytes"
)

type RLREApdu struct {
}

func (a RLREApdu) decode(is *bytes.Buffer, b bool) int {

}

func NewRLREApdu() *RLREApdu {
	return &RLREApdu{}
}
