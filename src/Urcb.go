package src

type Urcb struct {
	Rcb
	Test string
}

func NewUrcb(objectReference *ObjectReference, children []ModelNodeI) *Urcb {
	return &Urcb{Rcb: *NewRcb(objectReference, RP, children)}
}
