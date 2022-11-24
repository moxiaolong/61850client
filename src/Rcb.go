package src

import "unsafe"

type Rcb struct {
	FcDataObject
	dataSet *DataSet
}

func NewRcb(objectReference *ObjectReference, fc string, children []*FcModelNode) *Rcb {
	return &Rcb{FcDataObject: *NewFcDataObject(objectReference, fc, children)}
}

func (r *Rcb) getRptId() *BdaVisibleString {
	node := r.Children["RptID"]
	pointer := unsafe.Pointer(node)
	return (*BdaVisibleString)(pointer)
}
func (r *Rcb) getDatSet() *BdaVisibleString {
	node := r.Children["DatSet"]
	pointer := unsafe.Pointer(node)
	return (*BdaVisibleString)(pointer)
}
