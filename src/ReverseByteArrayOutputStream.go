package src

import (
	"bytes"
)

type ReverseByteArrayOutputStream struct {
	buffer []byte
	index  int
}

func (s *ReverseByteArrayOutputStream) getByteBuffer() *bytes.Buffer {
	return bytes.NewBuffer(s.buffer)
}

func (s *ReverseByteArrayOutputStream) getArray() []byte {
	if s.index == -1 {
		return s.buffer
	} else {
		subBuffer := s.buffer[s.index+1:]
		return subBuffer
	}
}

func (s *ReverseByteArrayOutputStream) reset() {
	s.index = len(s.buffer) - 1
}

func NewReverseByteArrayOutputStream(bufferSize int) *ReverseByteArrayOutputStream {
	if bufferSize <= 0 {
		Throw("buffer size may not be <= 0")
	}
	r := &ReverseByteArrayOutputStream{}
	r.buffer = make([]byte, bufferSize)
	r.index = bufferSize - 1
	r.index = bufferSize - 1
	return r
}

func (s *ReverseByteArrayOutputStream) write(byteArray []byte) {
	for {
		copy(s.buffer, byteArray)
		s.index -= len(byteArray)
		return
	}
}

func (s *ReverseByteArrayOutputStream) writeByte(byte byte) {
	for {
		//TODO
		s.buffer = append(s.buffer, byte)
		s.index -= 1
		return
	}
}

func (s *ReverseByteArrayOutputStream) read() int {
	return 0
}
