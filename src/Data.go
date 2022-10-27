package src

type Data struct {
	visibleString *BerVisibleString
	bitString     *BerBitString
	Unsigned      *BerInteger
	bool          *BerBoolean
	OctetString   *BerOctetString
}

func NewData() *Data {
	return &Data{}
}
