package src

type AssociateResult struct {
	BerInteger
}

func NewAssociateResult() *AssociateResult {
	return &AssociateResult{BerInteger: *NewBerInteger(nil, 0)}
}
