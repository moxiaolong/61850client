package src

type ServerModel struct {
	ModelNode
	urcbs   map[string]*Urcb
	brcbs   map[string]*Brcb
	dataSet map[string]*DataSet
}

func NewServerModel([]*LogicalDevice, []*DataSet) *ServerModel {
	return &ServerModel{ModelNode: *NewModelNode(), dataSet: make(map[string]*DataSet)}

}

func (m *ServerModel) getDataSet(ref string) *DataSet {
	return m.dataSet[ref]
}

func (m *ServerModel) getNodeFromVariableDef(def *SEQUENCE) *FcModelNode {
	ObjectName objectName = variableDef.getVariableSpecification().getName();

	if (objectName == nil) {
		throw (
			ServiceError.FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT,
			"name in objectName is not selected");
	}

	DomainSpecific domainSpecific = objectName.getDomainSpecific();

	if (domainSpecific == nil) {
		throw (
			ServiceError.FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT,
			"domain_specific in name is not selected");
	}

	ModelNode modelNode = getChild(domainSpecific.getDomainID().toString());

	if (modelNode == nil) {
		return nil;
	}

	String mmsItemId = domainSpecific.getItemID().toString();
	int index1 = mmsItemId.indexOf('$');

	if (index1 == -1) {
		throw (
			ServiceError.FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT,
			"invalid mms item id: " + domainSpecific.getItemID());
	}

	LogicalNode ln = (LogicalNode) modelNode.getChild(mmsItemId.substring(0, index1));

	if (ln == nil) {
		return nil;
	}

	int index2 = mmsItemId.indexOf('$', index1 + 1);

	if (index2 == -1) {
		throw (
			ServiceError.FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT, "invalid mms item id");
	}

	Fc fc = Fc.fromString(mmsItemId.substring(index1 + 1, index2));

	if (fc == nil) {
		throw (
			ServiceError.FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT,
			"unknown functional constraint: " + mmsItemId.substring(index1 + 1, index2));
	}

	index1 = index2;

	index2 = mmsItemId.indexOf('$', index1 + 1);

	if (index2 == -1) {
		if (fc == Fc.RP) {
			return ln.getUrcb(mmsItemId.substring(index1 + 1));
		}
		if (fc == Fc.BR) {
			return ln.getBrcb(mmsItemId.substring(index1 + 1));
		}
		return (FcModelNode) ln.getChild(mmsItemId.substring(index1 + 1), fc);
	}

	if (fc == Fc.RP) {
		modelNode = ln.getUrcb(mmsItemId.substring(index1 + 1, index2));
	} else if (fc == Fc.BR) {
		modelNode = ln.getBrcb(mmsItemId.substring(index1 + 1, index2));
	} else {
		modelNode = ln.getChild(mmsItemId.substring(index1 + 1, index2), fc);
	}

	index1 = index2;
	index2 = mmsItemId.indexOf('$', index1 + 1);
	while (index2 != -1) {
		modelNode = modelNode.getChild(mmsItemId.substring(index1 + 1, index2));
		index1 = index2;
		index2 = mmsItemId.indexOf('$', index1 + 1);
	}

	modelNode = modelNode.getChild(mmsItemId.substring(index1 + 1));

	if (variableDef.getAlternateAccess() == nil) {
		// no array is in this node path
		return (FcModelNode) modelNode;
	}

	AlternateAccessSelection altAccIt =
		variableDef.getAlternateAccess().getCHOICE().get(0).getUnnamed();

	if (altAccIt.getSelectAlternateAccess() != nil) {
		// path to node below an array element
		modelNode =
			((Array) modelNode)
		.getChild(
			altAccIt.getSelectAlternateAccess().getAccessSelection().getIndex().intValue());

		String mmsSubArrayItemId =
			altAccIt
		.getSelectAlternateAccess()
		.getAlternateAccess()
		.getCHOICE()
		.get(0)
		.getUnnamed()
		.getSelectAccess()
		.getComponent()
		.getBasic()
		.toString();

		index1 = -1;
		index2 = mmsSubArrayItemId.indexOf('$');
		while (index2 != -1) {
			modelNode = modelNode.getChild(mmsSubArrayItemId.substring(index1 + 1, index2));
			index1 = index2;
			index2 = mmsItemId.indexOf('$', index1 + 1);
		}

		return (FcModelNode) modelNode.getChild(mmsSubArrayItemId.substring(index1 + 1));
	} else {
		// path to an array element
		return (FcModelNode)
		((Array) modelNode).getChild(altAccIt.getSelectAccess().getIndex().intValue());
	}
}
}

func (m *ServerModel) addDataSet(set *DataSet) {
	dataSets.put(dataSet.getReferenceStr().replace('$', '.'), dataSet);
	for (ModelNode ld : children.values()) {
		for (ModelNode ln : ld.getChildren()) {
			for (Urcb urcb : ((LogicalNode) ln).getUrcbs()) {
				urcb.dataSet = getDataSet(urcb.getDatSet().getStringValue().replace('$', '.'));
			}
			for (Brcb brcb : ((LogicalNode) ln).getBrcbs()) {
				brcb.dataSet = getDataSet(brcb.getDatSet().getStringValue().replace('$', '.'));
			}
		}
	}
}

func (m *ServerModel) removeDataSet(ref string) {
	DataSet dataSet = dataSets.get(dataSetReference);
	if (dataSet == nil || !dataSet.isDeletable()) {
		return nil;
	}
	DataSet removedDataSet = dataSets.remove(dataSetReference);
	for (ModelNode ld : children.values()) {
		for (ModelNode ln : ld.getChildren()) {
			for (Urcb urcb : ((LogicalNode) ln).getUrcbs()) {
				urcb.dataSet = getDataSet(urcb.getDatSet().getStringValue().replace('$', '.'));
			}
			for (Brcb brcb : ((LogicalNode) ln).getBrcbs()) {
				brcb.dataSet = getDataSet(brcb.getDatSet().getStringValue().replace('$', '.'));
			}
		}
	}
	return removedDataSet;
}
