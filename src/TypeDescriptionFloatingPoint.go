package src

import "bytes"

type TypeDescriptionFloatingPoint struct {
	formatWidth *Unsigned8
}

func (p TypeDescriptionFloatingPoint) decode(is *bytes.Buffer, b bool) int {

}

func (p TypeDescriptionFloatingPoint) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewTypeDescriptionFloatingPoint() *TypeDescriptionFloatingPoint {
	return &TypeDescriptionFloatingPoint{}
}
