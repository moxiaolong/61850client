package src

type Integer8 struct {
	BerInteger
}

func NewInteger8(value int) *Integer8 {
	return &Integer8{BerInteger: *NewBerInteger(nil, value)}
}
