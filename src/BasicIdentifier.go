package src

type BasicIdentifier struct {
	BerVisibleString
}

func NewBasicIdentifier(value []byte) *BasicIdentifier {
	return &BasicIdentifier{BerVisibleString: *NewBerVisibleString(value)}
}
