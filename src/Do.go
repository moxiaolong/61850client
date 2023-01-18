package src

import "github.com/beevik/etree"

type Do struct {
	AbstractElement
	atype string
}

func (d *Do) getType() string {
	return d.atype
}
func NewDo(node *etree.Element) *Do {
	element := NewAbstractElement(node)
	do := &Do{AbstractElement: *element}

	atype := node.SelectAttr("type")

	do.atype = atype.Value
	return do
}
