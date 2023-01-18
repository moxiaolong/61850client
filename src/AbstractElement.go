package src

import "github.com/beevik/etree"

type AbstractElement struct {
	name string
	desc string
}

func (a *AbstractElement) getName() string {
	return a.name
}
func (a *AbstractElement) getDesc() string {
	return a.desc
}

func NewAbstractElement(node *etree.Element) *AbstractElement {
	a := &AbstractElement{}
	a.name = node.SelectAttr("name").Value
	attr := node.SelectAttr("desc")
	if attr != nil {
		a.desc = attr.Value
	}
	return a
}
