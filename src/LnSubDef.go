package src

import "github.com/beevik/etree"

type LnSubDef struct {
	logicalNode *LogicalNode
	defXmlNode  *etree.Element
}

func NewLnSubDef(defXmlNode *etree.Element, logicalNode *LogicalNode) *LnSubDef {
	return &LnSubDef{defXmlNode: defXmlNode, logicalNode: logicalNode}
}
