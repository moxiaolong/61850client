package src

type PDVList struct {
	PresentationContextIdentifier *PresentationContextIdentifier
	PresentationDataValues        *PDVListPresentationDataValues
}

func NewPDVList() *PDVList {
	return &PDVList{}
}
