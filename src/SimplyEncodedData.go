package src

type SimplyEncodedData struct {
	BerOctetString
}

func NewSimplyEncodedData() *SimplyEncodedData {
	return &SimplyEncodedData{BerOctetString: *NewBerOctetString(nil)}
}
