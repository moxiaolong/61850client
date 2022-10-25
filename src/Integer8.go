package src

type Integer8 struct {
}

func (i *Integer8) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0

}

func NewInteger8(int) *Integer8 {
	return &Integer8{}
}
