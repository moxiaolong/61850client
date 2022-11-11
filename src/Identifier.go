package src

type Identifier struct {
	BerVisibleString
}

func NewIdentifier() *Identifier {
	return &Identifier{BerVisibleString: *NewBerVisibleString()}
}
