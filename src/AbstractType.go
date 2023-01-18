package src

import "github.com/beevik/etree"

type AbstractType struct {
	Id string
}

func NewAbstractType(node *etree.Element) *AbstractType {

	id := node.SelectAttr("id").Value
	return &AbstractType{Id: id}
}
