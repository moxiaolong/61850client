package src

type ModelNode struct {
	children        map[string]*ModelNode
	objectReference *ObjectReference
}

func NewModelNode() *ModelNode {
	return &ModelNode{}
}
