package src

type Myexternal struct {
	DirectReference   *BerObjectIdentifier
	IndirectReference *BerInteger
	Encoding          *MyexternalEncoding
}

func (m *Myexternal) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewMyexternal() *Myexternal {
	return &Myexternal{}
}
