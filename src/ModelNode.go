package src

type ModelNodeI interface {
	getChild(name string, fc string) ModelNodeI
	getChildren() map[string]ModelNodeI
	getObjectReference() *ObjectReference
	setValueFromMmsDataObj(data *Data)
	getMmsDataObj() *Data
	copy() ModelNodeI
	setParent(node ModelNodeI)
	getName() string
	getParent() ModelNodeI
}

type ModelNode struct {
	Children        map[string]ModelNodeI
	ObjectReference *ObjectReference
	parent          ModelNodeI
}

func (m *ModelNode) copy() ModelNodeI {
	panic("impl me")
}
func (m *ModelNode) getMmsDataObj() *Data {
	panic("impl me")
}

func (m *ModelNode) setValueFromMmsDataObj(data *Data) {
	//none
}

func (m *ModelNode) setParent(node ModelNodeI) {
	m.parent = node
}

func (m *ModelNode) getParent() ModelNodeI {
	return m.parent
}

func (m *ModelNode) getChild(name string, fc string) ModelNodeI {
	return m.Children[name]
}
func (m *ModelNode) getChildren() map[string]ModelNodeI {
	return m.Children
}
func (m *ModelNode) getObjectReference() *ObjectReference {
	return m.ObjectReference
}

func (m *ModelNode) getName() string {
	return m.ObjectReference.getName()

}

func NewModelNode() *ModelNode {
	return &ModelNode{}
}
