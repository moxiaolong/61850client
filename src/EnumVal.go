package src

import (
	"github.com/beevik/etree"
	"strconv"
)

type EnumVal struct {
	id  string
	ord int
}

func (v *EnumVal) getId() string {
	return v.id
}

func (v *EnumVal) getOrd() int {
	return v.ord
}

func NewEnumVal(node *etree.Element) *EnumVal {
	e := &EnumVal{}
	e.id = node.Text()

	value := node.SelectAttr("ord").Value
	atoi, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	e.ord = atoi

	return e
}
