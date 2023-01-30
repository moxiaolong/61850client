package src

import (
	"strings"
)

type LogicalNode struct {
	ModelNode
	urcbs         map[string]*Urcb
	brcbs         map[string]*Brcb
	fcDataObjects map[string]map[string]FcDataObjectI
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
	dataObjectsCopy := make([]FcDataObjectI, 0)
	for _, obj := range n.Children {
		dataObjectsCopy = append(dataObjectsCopy, obj.copy().(FcDataObjectI))
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
			urcb.dataSet = (n.parent.getParent()).(*ServerModel).GetDataSet(strings.ReplaceAll(dataSetRef, "$", "."))
		}
	}
}

func NewLogicalNode(objectReference *ObjectReference, fcDataObjects []FcDataObjectI) *LogicalNode {
	l := &LogicalNode{}
	l.Children = make(map[string]ModelNodeI)
	l.fcDataObjects = make(map[string]map[string]FcDataObjectI)
	l.ObjectReference = objectReference
	l.urcbs = make(map[string]*Urcb)
	l.brcbs = make(map[string]*Brcb)

	for _, fcDataObject := range fcDataObjects {
		key := fcDataObject.GetObjectReference().getName() + fcDataObject.getFc()
		l.Children[key] = fcDataObject

		if l.fcDataObjects[fcDataObject.getFc()] == nil {
			l.fcDataObjects[fcDataObject.getFc()] = make(map[string]FcDataObjectI)
		}
		l.fcDataObjects[fcDataObject.getFc()][fcDataObject.GetObjectReference().getName()] = fcDataObject
		fcDataObject.setParent(l)
		if fcDataObject.getFc() == RP {
			l.addUrcb((fcDataObject).(*Urcb), false)
		} else if fcDataObject.getFc() == BR {
			l.addBrcb((fcDataObject).(*Brcb))
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
