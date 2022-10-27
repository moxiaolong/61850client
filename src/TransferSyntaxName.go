package src

type TransferSyntaxName struct {
	BerObjectIdentifier
}

func NewTransferSyntaxName() *TransferSyntaxName {
	return &TransferSyntaxName{BerObjectIdentifier: *NewBerObjectIdentifier(nil)}
}
