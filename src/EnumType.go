package src

import "github.com/beevik/etree"

type EnumType struct {
	AbstractType
	max    int
	min    int
	values []*EnumVal
}

func (t *EnumType) getValues() []*EnumVal {
	return t.values
}

func NewEnumType(node *etree.Element) *EnumType {
	abstractType := NewAbstractType(node)
	e := &EnumType{
		AbstractType: *abstractType,
	}

	for _, node := range node.ChildElements() {
		if node.Tag == ("EnumVal") {
			val := NewEnumVal(node)
			if val.getOrd() < e.min {
				e.min = val.getOrd()
			} else if val.getOrd() > e.max {
				e.max = val.getOrd()
			}
			e.values = append(e.values, val)
		}
	}

	return e
}
