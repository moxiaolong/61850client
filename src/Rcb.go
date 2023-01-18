package src

type RcbI interface {
	FcDataObjectI
	getRptId() *BdaVisibleString
	getDatSet() *BdaVisibleString
}
type Rcb struct {
	FcDataObject
	dataSet *DataSet
}

func NewRcb(objectReference *ObjectReference, fc string, children []ModelNodeI) *Rcb {
	return &Rcb{FcDataObject: *NewFcDataObject(objectReference, fc, children)}
}

func (r *Rcb) getRptId() *BdaVisibleString {
	node := r.Children["RptID"]
	return node.(*BdaVisibleString)
}
func (r *Rcb) getDatSet() *BdaVisibleString {
	node := r.Children["DatSet"]

	return node.(*BdaVisibleString)
}
