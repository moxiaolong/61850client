package src

type Unsigned32 struct {
	BerInteger
}

func NewUnsigned32(value int) *Unsigned32 {
	return &Unsigned32{BerInteger: *NewBerInteger(nil, value)}
}
