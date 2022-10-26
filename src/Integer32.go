package src

type Integer32 struct {
	BerInteger
}

func NewInteger32(value int) *Integer32 {
	return &Integer32{BerInteger: *NewBerInteger(nil, value)}
}
