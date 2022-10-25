package src

import "bytes"

type BerVisibleString struct {
}

func (s BerVisibleString) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func (s BerVisibleString) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0
}

func NewBerVisibleString() *BerVisibleString {
	return &BerVisibleString{}
}
