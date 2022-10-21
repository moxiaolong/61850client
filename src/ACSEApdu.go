package src

import (
	"bytes"
)

type ACSEApdu struct {
	Aarq *AARQApdu
	Aare *AAREApdu
}

func (a *ACSEApdu) encode(stream *ReverseByteArrayOutputStream) {

}

func (a *ACSEApdu) decode(ppdu *bytes.Buffer) {

}

func NewACSEApdu() *ACSEApdu {
	return &ACSEApdu{}
}
