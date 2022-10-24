package src

type AEQualifier struct {
	AeQualifierForm2 *AEQualifierForm2
}

func (q AEQualifier) encode(os *ReverseByteArrayOutputStream) int {
	//TODO
	return 0
}

func NewAEQualifier() *AEQualifier {
	return &AEQualifier{}
}
