package src

import (
	"bufio"
	"encoding/binary"
)

func readShort(r *bufio.Reader) int {
	shortBytes := make([]byte, 2)
	_, err := r.Read(shortBytes)
	if err != nil {
		panic(err)
	}
	return int(binary.BigEndian.Uint16(shortBytes))
}
func read(r *bufio.Reader) int {
	readByte, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	return int(readByte)
}

func readByte(r *bufio.Reader) byte {
	readByte, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	return readByte
}

func writeShort(w *bufio.Writer, v int) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(v))
	_, err := w.Write(buf)
	if err != nil {
		panic(err)
	}
}

func write(w *bufio.Writer, b int) {
	err := w.WriteByte(byte(b))
	if err != nil {
		panic(err)
	}
}

func writeByte(w *bufio.Writer, b byte) {
	err := w.WriteByte(b)
	if err != nil {
		panic(err)
	}
}
