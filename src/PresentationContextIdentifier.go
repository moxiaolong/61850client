package src

type PresentationContextIdentifier struct {
	BerInteger
}

func NewPresentationContextIdentifier(code []byte) *PresentationContextIdentifier {
	return &PresentationContextIdentifier{BerInteger: *NewBerInteger(code, 0)}
}
