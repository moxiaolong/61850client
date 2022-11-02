package src

type AEQualifierForm2 struct {
	BerInteger
}

func NewAEQualifierForm2(value int) *AEQualifierForm2 {
	return &AEQualifierForm2{BerInteger{Tag: NewBerTag(0, 0, 2), value: value}}
}
