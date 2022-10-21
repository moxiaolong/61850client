package src

import "bytes"

type ConcludeRequestPDU struct {
}

func (p ConcludeRequestPDU) decode(is *bytes.Buffer, b bool) int {
	return 0
}

func NewConcludeRequestPDU() *ConcludeRequestPDU {
	return &ConcludeRequestPDU{}

}
