package src

import "unsafe"

type ConstructedDataAttribute struct {
	FcModelNode
}

func NewConstructedDataAttribute(objectReference *ObjectReference, fc string, children []*FcModelNode) *ConstructedDataAttribute {
	c := &ConstructedDataAttribute{}
	c.ObjectReference = objectReference
	c.Fc = fc
	c.Children = make(map[string]*ModelNode)
	for _, child := range children {
		c.Children[child.getName()] = (*ModelNode)(unsafe.Pointer(child))
		child.parent = (*ModelNode)(unsafe.Pointer(c))
	}

	return c
}
