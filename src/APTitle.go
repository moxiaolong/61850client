package src

type APTitle struct {
	ApTitleForm2 *ApTitleForm2
}

func (t APTitle) encode(os *ReverseByteArrayOutputStream) int {
	//TODO
	return 1
}

func NewAPTitle() *APTitle {
	return &APTitle{}
}
