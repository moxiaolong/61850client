package src

import (
	"github.com/beevik/etree"
	"strconv"
)

type AbstractDataAttributeI interface {
	getbType() string
	getType() string
	getValue() string
	getCount() int
}

type AbstractDataAttribute struct {
	AbstractElement
	value string
	bType string
	atype string
	sAddr string
	count int
}

func NewAbstractDataAttribute(node *etree.Element) *AbstractDataAttribute {
	abstractElement := NewAbstractElement(node)
	attribute := &AbstractDataAttribute{AbstractElement: *abstractElement}

	for _, node := range node.Attr {
		nodeName := node.Key

		if nodeName == ("type") {
			attribute.atype = node.Value
		} else if nodeName == ("sAddr") {
			attribute.sAddr = node.Value
		} else if nodeName == ("bType") {
			attribute.bType = node.Value
		} else if nodeName == ("count") {
			atoi, err := strconv.Atoi(node.Value)
			if err != nil {
				panic(err)
			}
			attribute.count = atoi
		}
	}

	for _, node := range node.ChildElements() {
		if node.Tag == ("Val") {
			attribute.value = node.Text()
		}
	}

	return attribute
}

func (a *AbstractDataAttribute) getCount() int {
	return a.count

}

func (a *AbstractDataAttribute) getbType() string {
	return a.bType

}

func (a *AbstractDataAttribute) getType() string {
	return a.atype
}
func (a *AbstractDataAttribute) getValue() string {
	return a.value
}
