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

func (n *LogicalNode) addBrcb(brcb *Brcb) {
	n.brcbs[brcb.ObjectReference.getName()] = brcb
}

func (n *LogicalNode) addUrcb(urcb *Urcb, addDataSet bool) {

	n.urcbs[urcb.ObjectReference.getName()] = urcb
	if addDataSet {
		dataSetRef := urcb.getDatSet().getStringValue()
		if dataSetRef != "" {
			urcb.dataSet = (*ServerModel)(unsafe.Pointer(n.parent.parent)).getDataSet(strings.ReplaceAll(dataSetRef, "$", "."))
		}
	}
}

func NewLogicalNode(objectReference *ObjectReference, fcDataObjects []*FcDataObject) *LogicalNode {
	l := &LogicalNode{}
	l.Children = make(map[string]*ModelNode)
	l.fcDataObjects = make(map[string]map[string]*FcDataObject)
	l.ObjectReference = objectReference
	l.urcbs = make(map[string]*Urcb)
	l.brcbs = make(map[string]*Brcb)

	for _, fcDataObject := range fcDataObjects {
		key := fcDataObject.ObjectReference.getName() + fcDataObject.Fc
		l.Children[key] = (*ModelNode)(unsafe.Pointer(fcDataObject))

		if l.fcDataObjects[fcDataObject.Fc] == nil {
			l.fcDataObjects[fcDataObject.Fc] = make(map[string]*FcDataObject)
		}
		l.fcDataObjects[fcDataObject.Fc][fcDataObject.ObjectReference.getName()] = fcDataObject
		fcDataObject.parent = (*ModelNode)(unsafe.Pointer(l))
		if fcDataObject.Fc == RP {
			l.addUrcb((*Urcb)(unsafe.Pointer(fcDataObject)), false)
		} else if fcDataObject.Fc == BR {
			l.addBrcb((*Brcb)(unsafe.Pointer(fcDataObject)))
		}
	}
	return l
}

func (n *LogicalNode) getChild(childName string, fc string) *ModelNode {
	if fc != "" {
		return (*ModelNode)(unsafe.Pointer(n.fcDataObjects[fc][childName]))
	}
	for _, m := range n.fcDataObjects {
		fcDataObject := m[childName]
		if fcDataObject != nil {
			return (*ModelNode)(unsafe.Pointer(fcDataObject))
		}
	}

	return nil
}
