package src

type LogicalNode struct {
}

func NewLogicalNode(*ObjectReference, []*FcDataObject) *LogicalNode {
	return &LogicalNode{}
}
