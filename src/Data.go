package src

type Data struct {
	visibleString *BerVisibleString
	bitString     *BerBitString
	Unsigned      *BerInteger
}

func NewData() *Data {
	return &Data{}
}
