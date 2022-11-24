package src

import "unsafe"

type LogicalDevice struct {
	ModelNode
}

func NewLogicalDevice(objectReference *ObjectReference, logicalNodes []*LogicalNode) *LogicalDevice {
	node := &LogicalDevice{}
	node.Children = make(map[string]*ModelNode)
	node.ObjectReference = objectReference
	for _, logicalNode := range logicalNodes {
		node.Children[logicalNode.ObjectReference.getName()] = (*ModelNode)(unsafe.Pointer(logicalNode))
		logicalNode.parent = (*ModelNode)(unsafe.Pointer(node))
	}
	return node
}
