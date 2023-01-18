package src

import "github.com/beevik/etree"

type Bda struct {
	AbstractDataAttribute
}

func NewBda(node *etree.Element) *Bda {
	return &Bda{AbstractDataAttribute: *NewAbstractDataAttribute(node)}
}
