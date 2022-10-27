package src

import (
	"bytes"
)

type ReverseByteArrayOutputStream struct {
	buffer []byte
	index  int
}

func (s *ReverseByteArrayOutputStream) getByteBuffer() *bytes.Buffer {
	return bytes.NewBuffer(s.buffer[s.index+1:])
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
	return r
}

func (s *ReverseByteArrayOutputStream) write(byteArray []byte) {
	for {
		if s.index+1-len(byteArray) < 0 {
			s.resize()
			continue
		}

		//System.arraycopy(byteArray, 0, this.buffer, this.index-byteArray.length+1, byteArray.length)
		for i := len(byteArray) - 1; i >= 0; i-- {
			s.buffer[s.index] = byteArray[i]
			s.index -= 1
		}
		return
	}
}

func (s *ReverseByteArrayOutputStream) writeByte(byte byte) {
	defer func() {
		err := recover()
		if err != nil {
			s.resize()
			s.buffer[s.index] = byte
		}
	}()
	s.buffer[s.index] = byte
	s.index -= 1
	return

}

func (s *ReverseByteArrayOutputStream) resize() {
	newBuffer := make([]byte, len(s.buffer)*2)
	for i, b := range s.buffer {
		newBuffer[len(s.buffer)+i] = b
	}
	s.index += len(s.buffer)
	s.buffer = newBuffer
}
