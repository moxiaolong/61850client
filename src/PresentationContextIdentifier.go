package src

type PresentationContextIdentifier struct {
	BerInteger
}

func NewPresentationContextIdentifier(code []byte, value int) *PresentationContextIdentifier {
	return &PresentationContextIdentifier{BerInteger: *NewBerInteger(code, value)}
}
