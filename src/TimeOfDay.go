package src

type TimeOfDay struct {
	BerOctetString
}

func NewTimeOfDay(value []byte) *TimeOfDay {
	return &TimeOfDay{BerOctetString: *NewBerOctetString(value)}
}
