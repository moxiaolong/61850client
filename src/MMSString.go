package src

type MMSString struct {
	BerVisibleString
}

func NewMMSString() *MMSString {
	return &MMSString{BerVisibleString: *NewBerVisibleString()}
}
