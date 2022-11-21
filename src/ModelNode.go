package src

type ModelNode struct {
	children        map[string]*ModelNode
	objectReference *ObjectReference
	parent          *ModelNode
}

func (m *ModelNode) getChild(name string, fc string) *ModelNode {
	return m.children[name]
}

func (m *ModelNode) getName() string {
	return m.objectReference.getName()

}
func (m *ModelNode) copy() *ModelNode {
	//TODO implement me
	panic("implement me")
}

func NewModelNode() *ModelNode {
	return &ModelNode{}
}
