package src

type AEQualifier struct {
	AeQualifierForm2 *AEQualifierForm2
}

func (q *AEQualifier) encode(reverseOS *ReverseByteArrayOutputStream) int {

	codeLength := 0
	if q.AeQualifierForm2 != nil {
		codeLength += q.AeQualifierForm2.encode(reverseOS, true)
		return codeLength
	}

	Throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return -1
}

func NewAEQualifier() *AEQualifier {
	return &AEQualifier{}
}
