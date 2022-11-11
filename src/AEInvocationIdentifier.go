package src

type AEInvocationIdentifier struct {
	BerInteger
}

func NewAEInvocationIdentifier() *AEInvocationIdentifier {
	return &AEInvocationIdentifier{BerInteger: *NewBerInteger(nil, 0)}
}
