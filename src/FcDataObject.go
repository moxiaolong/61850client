package src

import (
	"strconv"
)

type FcDataObjectI interface {
	FcModelNodeI
	GetObjectReference() *ObjectReference
}

type FcDataObject struct {
	FcModelNode
}

func (n *FcDataObject) GetObjectReference() *ObjectReference {
	return n.ObjectReference
}
func (n *FcDataObject) getMmsDataObj() *Data {
	dataStructure := NewDataStructure()
	for _, modelNode := range n.getChildren() {
		child := modelNode.getMmsDataObj()
		if child == nil {
			throw("Unable to convert Child: " + modelNode.getObjectReference().toString() + " to MMS Data Object.")
		}
		dataStructure.seqOf = append(dataStructure.seqOf, child)
	}

	if len(dataStructure.seqOf) == 0 {
		throw("Converting ModelNode: " + n.ObjectReference.toString() + " to MMS Data Object resulted in Sequence of size zero.")
	}

	data := NewData()
	data.structure = dataStructure

	return data
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
