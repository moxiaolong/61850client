package src

type BasicIdentifier struct {
	BerVisibleString
}

func NewBasicIdentifier() *BasicIdentifier {
	return &BasicIdentifier{BerVisibleString: *NewBerVisibleString()}
}
