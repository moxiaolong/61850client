package src

import (
	"strings"
	"unsafe"
)

type ServerModel struct {
	ModelNode
	urcbs    map[string]*Urcb
	brcbs    map[string]*Brcb
	dataSets map[string]*DataSet
}

func NewServerModel([]*LogicalDevice, []*DataSet) *ServerModel {
	return &ServerModel{ModelNode: *NewModelNode(), dataSets: make(map[string]*DataSet)}

}

func (m *ServerModel) getDataSet(ref string) *DataSet {
	return m.dataSets[ref]
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

	modelNode := m.children[domainSpecific.domainID.toString()]

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
		return (*FcModelNode)(unsafe.Pointer(ln.getChild(mmsItemId[index1+1:0], fc)))
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
	m.dataSets[strings.ReplaceAll(dataSet.dataSetReference, "$", ".")] = dataSet

	for _, ld := range m.children {
		for _, ln := range ld.children {
			for _, urcb := range (*LogicalNode)(unsafe.Pointer(ln)).urcbs {
				urcb.dataSet = m.dataSets[(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))]
			}
			for _, brcb := range (*LogicalNode)(unsafe.Pointer(ln)).brcbs {
				brcb.dataSet = m.dataSets[(strings.ReplaceAll(brcb.getDatSet().getStringValue(), "$", "."))]
			}
		}
	}
}

func (m *ServerModel) removeDataSet(dataSetReference string) *DataSet {

	dataSet := m.dataSets[dataSetReference]
	if dataSet == nil || !dataSet.deletable {
		return nil
	}

	m.dataSets[dataSetReference] = nil
	removedDataSet := dataSet
	for _, ld := range m.children {
		for _, ln := range ld.children {
			for _, urcb := range (*LogicalNode)(unsafe.Pointer(ln)).urcbs {
				urcb.dataSet = m.dataSets[(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))]
			}
			for _, brcb := range (*LogicalNode)(unsafe.Pointer(ln)).brcbs {
				brcb.dataSet = m.dataSets[(strings.ReplaceAll(brcb.getDatSet().getStringValue(), "$", "."))]
			}
		}
	}

	return removedDataSet
}
