package src

type PresentationRequirements struct {
	BerBitString
}

func NewPresentationRequirements() *PresentationRequirements {
	return &PresentationRequirements{BerBitString: *NewBerBitString(nil, nil, 0)}
}
