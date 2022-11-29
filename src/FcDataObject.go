package src

import (
	"strconv"
)

type FcDataObject struct {
	FcModelNode
}

func (n *FcDataObject) setValueFromMmsDataObj(data *Data) {
	if data.structure == nil {
		throw("TYPE_CONFLICT expected type: structure")
	}
	if len(data.structure.seqOf) != len(n.Children) {
		throw(
			"TYPE_CONFLICT  expected type: structure with " + strconv.Itoa(len(n.Children)) + " elements")
	}
	index := 0
	for _, child := range n.Children {
		child.setValueFromMmsDataObj(data.structure.seqOf[index])
		index++
	}

}

func NewFcDataObject(objectReference *ObjectReference, fc string, children []ModelNodeI) *FcDataObject {
	f := &FcDataObject{}
	f.Children = make(map[string]ModelNodeI)
	f.ObjectReference = objectReference
	for _, child := range children {
		f.Children[child.getObjectReference().getName()] = (ModelNodeI)(child)
		child.setParent(f)
	}
	f.Fc = fc

	return f
}
