package src

type FcModelNode struct {
	ModelNode
	Fc string
}

func (n *FcModelNode) setValueFromMmsDataObj(success *Data) {

}

func (n *FcModelNode) copy() *FcModelNode {
	//TODO
	return n
}

func NewFcModelNode() *FcModelNode {

	return &FcModelNode{}
}
