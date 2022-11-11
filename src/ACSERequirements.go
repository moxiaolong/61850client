package src

type ACSERequirements struct {
	BerBitString
}

func NewACSERequirements() *ACSERequirements {
	return &ACSERequirements{BerBitString: *NewBerBitString(nil, nil, 0)}
}
