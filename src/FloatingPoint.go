package src

type FloatingPoint struct {
	BerOctetString
}

func NewFloatingPoint() *FloatingPoint {
	return &FloatingPoint{}
}
