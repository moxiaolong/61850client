package src

type FcDataObject struct {
	FcModelNode
}

func NewFcDataObject(*ObjectReference, string, []*FcModelNode) *FcDataObject {
	return &FcDataObject{FcModelNode: *NewFcModelNode()}
}
