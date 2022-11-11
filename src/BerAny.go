package src

import (
	"bytes"
)

type BerAny struct {
	value []byte
}

func (a *BerAny) encode(reverseOS *ReverseByteArrayOutputStream) int {
	reverseOS.write(a.value)
	return len(a.value)
}

func (a *BerAny) decode(is *bytes.Buffer, tag *BerTag) int {
	decodedLength := 0
	tagLength := 0
	if tag == nil {
		tag = NewEmptyBerTag()
		tagLength = tag.decode(is)
		decodedLength += tagLength
	} else {
		tagLength = tag.encode(NewReverseByteArrayOutputStream(10))
	}

	lengthField := NewBerLength()
	lengthLength := lengthField.decode(is)
	decodedLength += lengthLength + lengthField.val

	//跳过off 读l位
	off := tagLength + lengthLength
	l := lengthField.val
	a.value = make([]byte, l)
	_, err := is.Read(a.value)
	if err != nil {
		panic(err)
	}
	a.value = append(make([]byte, off), a.value...)

	os := NewReverseByteArrayOutputStreamWithBufferAndIndex(a.value, tagLength+lengthLength-1)
	encodeLength(os, lengthField.val)
	tag.encode(os)
	return decodedLength
}

func NewBerAny(value []byte) *BerAny {
	return &BerAny{value: value}
}
