package src

import "github.com/beevik/etree"

type DaType struct {
	AbstractType
	bdas []*Bda
}

func NewDaType(node *etree.Element) *DaType {
	abstractType := NewAbstractType(node)
	d := &DaType{AbstractType: *abstractType}

	for _, node := range node.ChildElements() {
		if node.Tag == "BDA" {
			d.bdas = append(d.bdas, NewBda(node))
		}
	}

	return d
}
