package src

type Myexternal struct {
	DirectReference   *BerObjectIdentifier
	IndirectReference *BerInteger
	Encoding          *MyexternalEncoding
}

func NewMyexternal() *Myexternal {
	return &Myexternal{}
}
