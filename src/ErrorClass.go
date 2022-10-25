package src

import (
	"bytes"
)

type ErrorClass struct {
}

func (c *ErrorClass) decode(is *bytes.Buffer) int {

	return 0
}

func (c *ErrorClass) encode(os *ReverseByteArrayOutputStream) int {
	return 0
}

func NewErrorClass() *ErrorClass {
	return &ErrorClass{}
}
