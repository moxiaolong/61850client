package src

type Brcb struct {
	Rcb
}

func NewBrcb(objectReference *ObjectReference, children []ModelNodeI) *Brcb {
	return &Brcb{Rcb: *NewRcb(objectReference, BR, children)}
}
