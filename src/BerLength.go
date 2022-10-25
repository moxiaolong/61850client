package src

import "bytes"

type BerLength struct {
	val int
}

func readEocByte(is *bytes.Buffer) int {
	return 0
}

func (l *BerLength) decode(buffer *bytes.Buffer) int {
	return 0
}

func (l *BerLength) readEocIfIndefinite(is *bytes.Buffer) int {
	return 0
}

func NewBerLength() *BerLength {
	return &BerLength{}
}
