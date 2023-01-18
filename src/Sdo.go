package src

import "github.com/beevik/etree"

type Sdo struct {
	AbstractElement
	atype string
}

func (s *Sdo) getType() string {
	return s.atype
}

func NewSdo(node *etree.Element) *Sdo {
	element := NewAbstractElement(node)
	sdo := &Sdo{AbstractElement: *element}

	attr := node.SelectAttr("type")
	sdo.atype = attr.Value
	return sdo
}
