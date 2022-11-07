package src

import "bytes"

type BerBoolean struct {
	value bool
}

func (b BerBoolean) decode(is *bytes.Buffer, b2 bool) int {

}

func (b BerBoolean) encode(os *ReverseByteArrayOutputStream, b2 bool) int {

}

func NewBerBoolean() *BerBoolean {
	return &BerBoolean{}
}
