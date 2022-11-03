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
		tag = NewBerTag(0, 0, 0)
		tagLength = tag.decode(is)
		decodedLength += tagLength
	} else {
		tagLength = tag.encode(NewReverseByteArrayOutputStream(10))
	}

	lengthField := NewBerLength()
	lengthLength := lengthField.decode(is)
	decodedLength += lengthLength + lengthField.val
	a.value = make([]byte, tagLength+lengthLength+lengthField.val)
	_, err := is.Read(a.value)
	if err != nil {
		panic(err)
	}
	os := NewReverseByteArrayOutputStreamWithBufferAndIndex(a.value, tagLength+lengthLength-1)
	encodeLength(os, lengthField.val)
	tag.encode(os)
	return decodedLength
}

func NewBerAny(value []byte) *BerAny {
	return &BerAny{value: value}
}
