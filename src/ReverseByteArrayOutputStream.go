package src

import "bytes"

type ReverseByteArrayOutputStream struct {
}

func (s *ReverseByteArrayOutputStream) getByteBuffer() *bytes.Buffer {
	return &bytes.Buffer{}
}

func NewReverseByteArrayOutputStream(i int, b bool) *ReverseByteArrayOutputStream {
	return nil
}
