package src

import "unsafe"

type Rcb struct {
	FcDataObject
}

func NewRcb() *Rcb {
	return &Rcb{}
}

func (r *Rcb) getRptId() *BdaVisibleString {
	node := r.children["RptID"]
	pointer := unsafe.Pointer(node)
	return (*BdaVisibleString)(pointer)
}
func (r *Rcb) getDatSet() *BdaVisibleString {
	node := r.children["DatSet"]
	pointer := unsafe.Pointer(node)
	return (*BdaVisibleString)(pointer)
}