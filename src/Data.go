package src

type Data struct {
	visibleString *BerVisibleString
	bitString     *BerBitString
	Unsigned      *BerInteger
	bool          *BerBoolean
	OctetString   *BerOctetString
	binaryTime    *TimeOfDay
}

func NewData() *Data {
	return &Data{}
}
