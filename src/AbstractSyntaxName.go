package src

type AbstractSyntaxName struct {
	BerObjectIdentifier
}

func NewAbstractSyntaxName() *AbstractSyntaxName {
	return &AbstractSyntaxName{BerObjectIdentifier: *NewBerObjectIdentifier(nil)}
}
