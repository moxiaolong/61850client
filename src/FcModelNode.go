package src

type FcModelNode struct {
	ModelNode
	Fc string
}

func (n *FcModelNode) setValueFromMmsDataObj(success *Data) {

}

func NewFcModelNode() *FcModelNode {
	return &FcModelNode{ModelNode: *NewModelNode()}
}
