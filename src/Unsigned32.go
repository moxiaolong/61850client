package src

type Unsigned32 struct {
	BerInteger
}

func NewUnsigned32() *Unsigned32 {
	return &Unsigned32{}
}
