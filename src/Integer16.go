package src

type Integer16 struct {
	BerInteger
}

func NewInteger16(code []byte, value int) *Integer16 {
	return &Integer16{BerInteger: *NewBerInteger(code, value)}
}
