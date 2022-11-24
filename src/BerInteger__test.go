package src

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func Test(t *testing.T) {
	integer := NewBerInteger(nil, 0)
	buf := bytes.NewBuffer([]byte{1, 0xF3})
	integer.decode(buf, false)

	println(integer.intValue())

}

func TestName(t *testing.T) {
	data := int64(1)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	print(bytebuf.Bytes())
}
