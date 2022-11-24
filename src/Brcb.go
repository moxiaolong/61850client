package src

type Brcb struct {
	Rcb
}

func NewBrcb(objectReference *ObjectReference, children []*FcModelNode) *Brcb {
	return &Brcb{Rcb: *NewRcb(objectReference, BR, children)}
}
