package src

type PDVListPresentationDataValues struct {
	SingleASN1Type *BerAny
}

func NewPDVListPresentationDataValues() *PDVListPresentationDataValues {
	return &PDVListPresentationDataValues{}
}
