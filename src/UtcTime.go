package src

type UtcTime struct {
	BerOctetString
}

func NewUtcTime() *UtcTime {
	return &UtcTime{BerOctetString: *NewBerOctetString(nil)}
}
