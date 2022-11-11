package src

type ImplementationData struct {
	BerGraphicString
}

func NewImplementationData() *ImplementationData {
	return &ImplementationData{BerGraphicString: *NewBerGraphicString()}
}
