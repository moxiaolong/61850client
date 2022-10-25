package src

type Integer32 struct {
}

func (i *Integer32) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0

}

func NewInteger32(int) *Integer32 {
	return &Integer32{}
}
