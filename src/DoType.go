package src

import "github.com/beevik/etree"

type DoType struct {
	AbstractType
	das  []*Da
	sdos []*Sdo
}

func NewDoType(node *etree.Element) *DoType {
	abstractType := NewAbstractType(node)
	doType := &DoType{
		AbstractType: *abstractType,
	}

	if node.SelectAttr("cdc") == nil {
		panic("Required attribute \"cdc\" not found in DOType!")
	}

	for _, node := range node.ChildElements() {
		if node.Tag == "SDO" {
			doType.sdos = append(doType.sdos, NewSdo(node))
		}
		if node.Tag == ("DA") {
			doType.das = append(doType.das, NewDa(node))
		}
	}

	return doType
}
