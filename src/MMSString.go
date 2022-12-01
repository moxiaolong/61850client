package src

type MMSString struct {
	BerVisibleString
}

func NewMMSString(value []byte) *MMSString {
	return &MMSString{BerVisibleString: *NewBerVisibleString(value)}
}
