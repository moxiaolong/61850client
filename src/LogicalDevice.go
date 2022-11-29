package src

type LogicalDevice struct {
	ModelNode
}

func NewLogicalDevice(objectReference *ObjectReference, logicalNodes []ModelNodeI) *LogicalDevice {
	node := &LogicalDevice{}
	node.Children = make(map[string]ModelNodeI)
	node.ObjectReference = objectReference
	for _, logicalNode := range logicalNodes {
		node.Children[logicalNode.getObjectReference().getName()] = logicalNode
		logicalNode.setParent(node)
	}
	return node
}
