package src

import "bytes"

type ReverseByteArrayOutputStream struct {
	buffer []byte
	index  int
}

func (s *ReverseByteArrayOutputStream) getByteBuffer() *bytes.Buffer {
	return &bytes.Buffer{}
}

func (s *ReverseByteArrayOutputStream) getArray() []byte {
	return nil
}

func (s *ReverseByteArrayOutputStream) reset() {

}

func NewReverseByteArrayOutputStream(i int, b bool) *ReverseByteArrayOutputStream {
	return &ReverseByteArrayOutputStream{}
}
