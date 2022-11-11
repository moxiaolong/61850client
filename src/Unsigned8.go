package src

type Unsigned8 struct {
	BerInteger
}

func NewUnsigned8() *Unsigned8 {
	return &Unsigned8{BerInteger: *NewBerInteger(nil, 0)}
}
