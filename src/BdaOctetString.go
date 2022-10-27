package src

type BdaOctetString struct {
}

func (s BdaOctetString) setValue(value []byte) {

}

func NewBdaOctetString(*ObjectReference, interface{}, string, int, bool, bool) *BdaOctetString {
	return &BdaOctetString{}
}
