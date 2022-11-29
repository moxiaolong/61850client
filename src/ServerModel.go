package src

import (
	"strings"
)

type ServerModel struct {
	ModelNode
	urcbs    map[string]*Urcb
	brcbs    map[string]*Brcb
	DataSets map[string]*DataSet
}

func NewServerModel(logicalDevices []*LogicalDevice, dataSets []*DataSet) *ServerModel {
	m := &ServerModel{}

	m.Children = make(map[string]ModelNodeI)
	m.ObjectReference = nil
	m.urcbs = make(map[string]*Urcb)
	m.brcbs = make(map[string]*Brcb)
	m.DataSets = make(map[string]*DataSet)

	for _, logicalDevice := range logicalDevices {
		m.Children[logicalDevice.ObjectReference.getName()] = (ModelNodeI)(logicalDevice)
		logicalDevice.parent = m
	}

	m.addDataSets(dataSets)

	for _, ld := range logicalDevices {
		for _, ln := range ld.Children {
			l := (ln).(*LogicalNode)
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

func (m *ServerModel) getNodeFromVariableDef(variableDef *VariableDefsSEQUENCE) ModelNodeI {
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

	//ln := (*LogicalNode)(modelNode.getChild(mmsItemId[0:index1], ""))
	ln := (modelNode.getChild(mmsItemId[0:index1], "")).(*LogicalNode)

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
			return urcb
		}
		if fc == BR {
			brcb := ln.brcbs[mmsItemId[index1+1:]]
			return brcb
		}
		child := ln.getChild(mmsItemId[index1+1:], fc)
		return child
	}
	index2 += index1 + 1

	if fc == RP {
		urcb := ln.urcbs[mmsItemId[index1+1:index2]]
		modelNode = urcb
	} else if fc == BR {
		brcb := ln.brcbs[mmsItemId[index1+1:index2]]
		modelNode = brcb
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
		return modelNode
	}

	altAccIt :=
		variableDef.alternateAccess.seqOf[0].unnamed

	if altAccIt.selectAlternateAccess != nil {
		// path to node below an array element
		modelNode =
			(modelNode.(*FCArray)).getChildIndex(altAccIt.selectAlternateAccess.accessSelection.index.intValue())

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
		return child
	} else {
		// path to an array element
		node := modelNode.(*FCArray).getChildIndex(altAccIt.selectAccess.index.intValue())
		return node
	}
}

func (m *ServerModel) addDataSet(dataSet *DataSet) {
	m.DataSets[strings.ReplaceAll(dataSet.DataSetReference, "$", ".")] = dataSet

	for _, ld := range m.Children {
		for _, ln := range ld.getChildren() {
			for _, urcb := range ln.(*LogicalNode).urcbs {
				urcb.dataSet = m.DataSets[(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))]
			}
			for _, brcb := range ln.(*LogicalNode).brcbs {
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
		for _, ln := range ld.getChildren() {
			for _, urcb := range ln.(*LogicalNode).urcbs {
				urcb.dataSet = m.DataSets[(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))]
			}
			for _, brcb := range ln.(*LogicalNode).brcbs {
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
		for _, ln := range ld.getChildren() {
			l := ln.(*LogicalNode)
			for _, urcb := range l.urcbs {
				urcb.dataSet = m.getDataSet(strings.ReplaceAll(urcb.getDatSet().getStringValue(), "$", "."))
			}
			for _, brcb := range l.brcbs {
				brcb.dataSet = m.getDataSet(strings.ReplaceAll(brcb.getDatSet().getStringValue(), "$", "."))
			}
		}
	}

}

func (m *ServerModel) findModelNode(objectReferenceStr string, fc string) ModelNodeI {
	currentNode := (ModelNodeI)(m)
	objectReference := NewObjectReference(objectReferenceStr)
	objectReference.parseForNameList()
	for _, name := range objectReference.nodeNames {
		currentNode = currentNode.getChild(name, fc)
		if currentNode == nil {
			return nil
		}

	}
	return currentNode
}

func (m *ServerModel) AskForFcModelNode(objectReferenceStr string, fc string) FcModelNodeI {
	modelNode := m.findModelNode(objectReferenceStr, fc)
	if modelNode == nil {
		throw("A model node with the given reference and functional constraint could not be found.")
	}

	return modelNode.(FcModelNodeI)

}
