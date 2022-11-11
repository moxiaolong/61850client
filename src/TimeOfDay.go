package src

type TimeOfDay struct {
	BerOctetString
}

func NewTimeOfDay() *TimeOfDay {
	return &TimeOfDay{BerOctetString: *NewBerOctetString(nil)}
}
