package src

type BdaVisibleStringI interface {
}
type BdaVisibleString struct {
	BasicDataAttribute
}

func (s *BdaVisibleString) getStringValue() string {

}

func NewBdaVisibleString() *BdaVisibleString {
	return &BdaVisibleString{}
}
