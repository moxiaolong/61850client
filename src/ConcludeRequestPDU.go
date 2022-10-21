package src

type ConcludeRequestPDU struct {
}

func (p ConcludeRequestPDU) decode(is *ByteBufferInputStream, b bool) int {
	return 0
}

func NewConcludeRequestPDU() *ConcludeRequestPDU {
	return &ConcludeRequestPDU{}

}
