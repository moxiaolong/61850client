package src

type Urcb struct {
	Rcb
	Test string
}

func NewUrcb(objectReference *ObjectReference, children []*FcModelNode) *Urcb {
	return &Urcb{Rcb: *NewRcb(objectReference, RP, children)}
}
