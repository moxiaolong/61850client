package src

type ConcludeRequestPDU struct {
	BerNull
}

func NewConcludeRequestPDU() *ConcludeRequestPDU {
	berNull := NewBerNull()
	return &ConcludeRequestPDU{*berNull}

}
