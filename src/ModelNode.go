package src

type ModelNode struct {
	children        map[string]*ModelNode
	objectReference *ObjectReference
}

func NewModelNode() *ModelNode {
	return &ModelNode{}
}

func (n *ModelNode) copy() *ModelNode {

}
