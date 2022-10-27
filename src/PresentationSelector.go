package src

type PresentationSelector struct {
	BerOctetString
}

func NewPresentationSelector(value []byte) *PresentationSelector {
	return &PresentationSelector{BerOctetString: *NewBerOctetString(value)}
}
