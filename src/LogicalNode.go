package src

type LogicalNode struct {
	ModelNode
	urcbs map[string]*Urcb
	brcbs map[string]*Brcb
}

func NewLogicalNode(*ObjectReference, []*FcDataObject) *LogicalNode {
	return &LogicalNode{}
}
