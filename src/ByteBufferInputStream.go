package src

import "bytes"

type ByteBufferInputStream struct {
	Buf *bytes.Buffer
}

func (s *ByteBufferInputStream) read() int {

	if s.Buf.Len() < 0 {
		return -1
	}
	readByte, err := s.Buf.ReadByte()
	if err != nil {
		return -1
	}
	return int(readByte & 0xFF)
}

func NewByteBufferInputStream(buf *bytes.Buffer) *ByteBufferInputStream {
	return &ByteBufferInputStream{Buf: buf}
}
