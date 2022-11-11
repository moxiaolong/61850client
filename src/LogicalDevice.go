package src

type LogicalDevice struct {
	ModelNode
}

func NewLogicalDevice(*ObjectReference, []*LogicalNode) *LogicalDevice {
	return &LogicalDevice{ModelNode: *NewModelNode()}
}
