package src

type APInvocationIdentifier struct {
	BerInteger
}

func NewAPInvocationIdentifier() *APInvocationIdentifier {
	return &APInvocationIdentifier{}
}
