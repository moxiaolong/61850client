package src

type Brcb struct {
	Rcb
}

func NewBrcb(*ObjectReference, []*FcModelNode) *Brcb {
	return &Brcb{}
}
