package src

type ModelNode struct {
	Children        map[string]*ModelNode
	ObjectReference *ObjectReference
	parent          *ModelNode
}

func (m *ModelNode) getChild(name string, fc string) *ModelNode {
	return m.Children[name]
}

func (m *ModelNode) getName() string {
	return m.ObjectReference.getName()

}
func (m *ModelNode) copy() *ModelNode {
	//TODO implement me
	panic("implement me")
}

func NewModelNode() *ModelNode {
	return &ModelNode{}
}
