package src

import (
	"bytes"
	"encoding/binary"
)

type BerInteger struct {
	code  []byte
	value int
	Tag   *BerTag
}

func (t *BerInteger) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if t.code != nil {
		reverseOS.write(t.code)
		if withTag {
			return t.Tag.encode(reverseOS) + len(t.code)
		} else {
			return len(t.code)
		}

	} else {
		buffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(buffer, binary.BigEndian, int64(t.value))
		//仿照Java bigint 转 byteArray 大端 保留有效位
		var buf []byte
		for buffer.Len() > 0 {
			b, _ := buffer.ReadByte()
			if b != 0 {
				buf = append([]byte{b}, buffer.Bytes()...)
				break
			}
		}
		if buf == nil {
			buf = []byte{0}
		}
		//补正负位
		heightBit := buf[0] >> 7
		if heightBit == 1 {
			if t.value < 0 {
				buf = append([]byte{1}, buf...)
			} else {
				buf = append([]byte{0}, buf...)
			}
		}

		codeLength := len(buf)
		reverseOS.write(buf)
		codeLength += encodeLength(reverseOS, codeLength)
		if withTag {
			codeLength += t.Tag.encode(reverseOS)
		}

		return codeLength
	}
}

func (t *BerInteger) decode(is *bytes.Buffer, withTag bool) int {
	codeLength := 0
	if withTag {
		codeLength += t.Tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	codeLength += length.decode(is)
	if length.val < 1 {
		throw("Decoded length of BerInteger is not correct")
	} else {

		byteCode := make([]byte, length.val)
		readFully(is, byteCode)
		codeLength += length.val
		//TODO
		t.value = int(binary.LittleEndian.Uint64(byteCode))
		return codeLength
	}
	return -1
}

func NewBerInteger(code []byte, value int) *BerInteger {
	return &BerInteger{code: code, value: value, Tag: NewBerTag(0, 0, 2)}
}
