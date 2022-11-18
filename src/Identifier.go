package src

type Identifier struct {
	BerVisibleString
}

func NewIdentifier(value []byte) *Identifier {
	return &Identifier{BerVisibleString: *NewBerVisibleString(value)}
}
