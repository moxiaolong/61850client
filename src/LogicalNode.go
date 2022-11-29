package src

import (
	"strings"
	"unsafe"
)

type LogicalNode struct {
	ModelNode
	urcbs         map[string]*Urcb
	brcbs         map[string]*Brcb
	fcDataObjects map[string]map[string]*FcDataObject
}

func (n *LogicalNode) setValueFromMmsDataObj(data *Data) {
	//TODO implement me
	panic("implement me")
}

func (n *LogicalNode) getMmsVariableDef() *VariableDefsSEQUENCE {
	//TODO implement me
	panic("implement me")
}

func (n *LogicalNode) copy() ModelNodeI {
	dataObjectsCopy := make([]*FcDataObject, 0)
	for _, obj := range n.Children {
		dataObjectsCopy = append(dataObjectsCopy, obj.copy().(*FcDataObject))
	}
	newCopy := NewLogicalNode(n.ObjectReference, dataObjectsCopy)
	return newCopy
}

func (n *LogicalNode) addBrcb(brcb *Brcb) {
	n.brcbs[brcb.ObjectReference.getName()] = brcb
}

func (n *LogicalNode) addUrcb(urcb *Urcb, addDataSet bool) {

	n.urcbs[urcb.ObjectReference.getName()] = urcb
	if addDataSet {
		dataSetRef := urcb.getDatSet().getStringValue()
		if dataSetRef != "" {
			urcb.dataSet = (n.parent.getParent()).(*ServerModel).getDataSet(strings.ReplaceAll(dataSetRef, "$", "."))
		}
	}
}

func NewLogicalNode(objectReference *ObjectReference, fcDataObjects []*FcDataObject) *LogicalNode {
	l := &LogicalNode{}
	l.Children = make(map[string]ModelNodeI)
	l.fcDataObjects = make(map[string]map[string]*FcDataObject)
	l.ObjectReference = objectReference
	l.urcbs = make(map[string]*Urcb)
	l.brcbs = make(map[string]*Brcb)

	for _, fcDataObject := range fcDataObjects {
		key := fcDataObject.ObjectReference.getName() + fcDataObject.Fc
		l.Children[key] = fcDataObject

		if l.fcDataObjects[fcDataObject.Fc] == nil {
			l.fcDataObjects[fcDataObject.Fc] = make(map[string]*FcDataObject)
		}
		l.fcDataObjects[fcDataObject.Fc][fcDataObject.ObjectReference.getName()] = fcDataObject
		fcDataObject.parent = l
		if fcDataObject.Fc == RP {
			l.addUrcb((*Urcb)(unsafe.Pointer(fcDataObject)), false)
		} else if fcDataObject.Fc == BR {
			l.addBrcb((*Brcb)(unsafe.Pointer(fcDataObject)))
		}
	}
	return l
}

func (n *LogicalNode) getChild(childName string, fc string) ModelNodeI {
	if fc != "" {
		object := n.fcDataObjects[fc][childName]
		return (ModelNodeI)(object)
	}
	for _, m := range n.fcDataObjects {
		fcDataObject := m[childName]
		if fcDataObject != nil {
			return fcDataObject
		}
	}

	return nil
}
