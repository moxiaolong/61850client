package src

import (
	"strings"
	"unsafe"
)

type ServerModel struct {
	ModelNode
	urcbs    map[string]*Urcb
	brcbs    map[string]*Brcb
	DataSets map[string]*DataSet
}

func NewServerModel(logicalDevices []*LogicalDevice, dataSets []*DataSet) *ServerModel {
	m := &ServerModel{}

	m.Children = make(map[string]*ModelNode)
	m.ObjectReference = nil
	m.urcbs = make(map[string]*Urcb)
	m.brcbs = make(map[string]*Brcb)
	m.DataSets = make(map[string]*DataSet)

	for _, logicalDevice := range logicalDevices {
		m.Children[logicalDevice.ObjectReference.getName()] = (*ModelNode)(unsafe.Pointer(logicalDevice))
		logicalDevice.parent = (*ModelNode)(unsafe.Pointer(m))
	}

	m.addDataSets(dataSets)

	for _, ld := range logicalDevices {
		for _, ln := range ld.Children {
			l := (*LogicalNode)(unsafe.Pointer(ln))
			for _, urcb := range l.urcbs {
				m.urcbs[urcb.ObjectReference.toString()] = urcb

				urcb.dataSet = m.getDataSet(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))

			}
			for _, brcb := range l.brcbs {
				m.brcbs[brcb.ObjectReference.toString()] = brcb
				brcb.dataSet = m.getDataSet(strings.ReplaceAll(brcb.getDatSet().getStringValue(), "$", "."))
			}
		}
	}

	return m

}

func (m *ServerModel) getDataSet(ref string) *DataSet {
	return m.DataSets[ref]
}

func (m *ServerModel) getNodeFromVariableDef(variableDef *VariableDefsSEQUENCE) *FcModelNode {
	objectName := variableDef.variableSpecification.name

	if objectName == nil {
		throw(
			"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT name in objectName is not selected")
	}

	domainSpecific := objectName.domainSpecific

	if domainSpecific == nil {
		throw(
			"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT domain_specific in name is not selected")
	}

	modelNode := m.Children[domainSpecific.domainID.toString()]

	if modelNode == nil {
		return nil
	}

	mmsItemId := domainSpecific.itemID.toString()

	index1 := strings.Index(mmsItemId, "$")

	if index1 == -1 {
		throw(
			"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT invalid mms item id: " + domainSpecific.itemID.toString())
	}

	ln := (*LogicalNode)(unsafe.Pointer(modelNode.getChild(mmsItemId[0:index1], "")))

	if ln == nil {
		return nil
	}

	//index2 := strings.Index("$") mmsItemId.indexOf('$', index1+1)
	index2 := strings.Index(mmsItemId[index1+1:], "$")
	if index2 == -1 {
		throw(
			"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT invalid mms item id")
	}
	index2 += index1 + 1

	fc := mmsItemId[index1+1 : index2]

	if fc == "" {
		throw(
			"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT unknown functional constraint: " + mmsItemId[index1+1:index2])
	}

	index1 = index2

	index2 = strings.Index(mmsItemId[index1+1:], "$")

	if index2 == -1 {
		if fc == RP {
			urcb := ln.urcbs[(mmsItemId[index1+1:])]
			return (*FcModelNode)(unsafe.Pointer(urcb))
		}
		if fc == BR {
			brcb := ln.brcbs[mmsItemId[index1+1:]]
			return (*FcModelNode)(unsafe.Pointer(brcb))
		}
		return (*FcModelNode)(unsafe.Pointer(ln.getChild(mmsItemId[index1+1:], fc)))
	}
	index2 += index1 + 1

	if fc == RP {
		urcb := ln.urcbs[mmsItemId[index1+1:index2]]
		modelNode = (*ModelNode)(unsafe.Pointer(urcb))
	} else if fc == BR {
		brcb := ln.brcbs[mmsItemId[index1+1:index2]]
		modelNode = (*ModelNode)(unsafe.Pointer(brcb))
	} else {
		modelNode = ln.getChild(mmsItemId[index1+1:index2], fc)
	}

	index1 = index2
	index2 = strings.Index(mmsItemId[index1+1:], "$")
	if index2 != -1 {
		index2 += index1 + 1
	}
	for index2 != -1 {
		modelNode = modelNode.getChild(mmsItemId[index1+1:index2], "")
		index1 = index2
		index2 = strings.Index(mmsItemId[index1+1:], "$")
		if index2 != -1 {
			index2 += index1 + 1
		}
	}

	modelNode = modelNode.getChild(mmsItemId[index1+1:], "")

	if variableDef.alternateAccess == nil {
		// no array is in this node path
		return (*FcModelNode)(unsafe.Pointer(modelNode))
	}

	altAccIt :=
		variableDef.alternateAccess.seqOf[0].unnamed

	if altAccIt.selectAlternateAccess != nil {
		// path to node below an array element
		modelNode =
			((*Array)(unsafe.Pointer(modelNode))).getChildIndex(altAccIt.selectAlternateAccess.accessSelection.index.intValue())

		mmsSubArrayItemId :=
			altAccIt.selectAlternateAccess.alternateAccess.seqOf[0].unnamed.selectAccess.component.basic.toString()

		index1 = -1
		index2 = strings.Index(mmsSubArrayItemId, "$")
		for index2 != -1 {
			modelNode = modelNode.getChild(mmsSubArrayItemId[index1+1:index2], "")
			index1 = index2
			index2 = strings.Index(mmsItemId[index1+1:], "$")
			if index2 != -1 {
				index2 += index1 + 1
			}
		}

		child := modelNode.getChild(mmsSubArrayItemId[index1:1], "")
		return (*FcModelNode)(unsafe.Pointer(child))
	} else {
		// path to an array element
		node := (*Array)(unsafe.Pointer(modelNode)).getChildIndex(altAccIt.selectAccess.index.intValue())
		return (*FcModelNode)(unsafe.Pointer(node))
	}
}

func (m *ServerModel) addDataSet(dataSet *DataSet) {
	m.DataSets[strings.ReplaceAll(dataSet.DataSetReference, "$", ".")] = dataSet

	for _, ld := range m.Children {
		for _, ln := range ld.Children {
			for _, urcb := range (*LogicalNode)(unsafe.Pointer(ln)).urcbs {
				urcb.dataSet = m.DataSets[(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))]
			}
			for _, brcb := range (*LogicalNode)(unsafe.Pointer(ln)).brcbs {
				brcb.dataSet = m.DataSets[(strings.ReplaceAll(brcb.getDatSet().getStringValue(), "$", "."))]
			}
		}
	}
}

func (m *ServerModel) removeDataSet(dataSetReference string) *DataSet {

	dataSet := m.DataSets[dataSetReference]
	if dataSet == nil || !dataSet.deletable {
		return nil
	}

	m.DataSets[dataSetReference] = nil
	removedDataSet := dataSet
	for _, ld := range m.Children {
		for _, ln := range ld.Children {
			for _, urcb := range (*LogicalNode)(unsafe.Pointer(ln)).urcbs {
				urcb.dataSet = m.DataSets[(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))]
			}
			for _, brcb := range (*LogicalNode)(unsafe.Pointer(ln)).brcbs {
				brcb.dataSet = m.DataSets[(strings.ReplaceAll(brcb.getDatSet().getStringValue(), "$", "."))]
			}
		}
	}

	return removedDataSet
}

func (m *ServerModel) addDataSets(dataSets []*DataSet) {
	for _, dataSet := range dataSets {
		m.addDataSet(dataSet)
	}
	for _, ld := range m.Children {
		for _, ln := range ld.Children {
			l := (*LogicalNode)(unsafe.Pointer(ln))
			for _, urcb := range l.urcbs {
				urcb.dataSet = m.getDataSet(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))
			}
			for _, brcb := range l.brcbs {
				brcb.dataSet = m.getDataSet(strings.ReplaceAll(brcb.getDatSet().getStringValue(), "$", "."))
			}
		}
	}

}
