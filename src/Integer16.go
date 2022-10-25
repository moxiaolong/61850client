package src

type Integer16 struct {
}

func (i *Integer16) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0

}

func NewInteger16([]byte) *Integer16 {
	return &Integer16{}
}
func NewInteger16Int(int) *Integer16 {
	return &Integer16{}
}
