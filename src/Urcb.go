package src

type Urcb struct {
	Rcb
	Test string
}

func NewUrcb(*ObjectReference, []*FcModelNode) *Urcb {
	return &Urcb{}
}
