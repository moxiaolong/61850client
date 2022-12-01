package src

type FloatingPoint struct {
	BerOctetString
}

func NewFloatingPoint(value []byte) *FloatingPoint {
	return &FloatingPoint{BerOctetString: *NewBerOctetString(value)}
}
