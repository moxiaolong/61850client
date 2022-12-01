package src

type UtcTime struct {
	BerOctetString
}

func NewUtcTime(value []byte) *UtcTime {
	return &UtcTime{BerOctetString: *NewBerOctetString(value)}
}
