package src

import "bytes"

type VariableDefs struct {
}

func (d VariableDefs) decode(is *bytes.Buffer, b bool) int {

}

func (d VariableDefs) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewVariableDefs() *VariableDefs {
	return &VariableDefs{}
}
