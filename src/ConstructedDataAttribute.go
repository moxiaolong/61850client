package src

import "unsafe"

type ConstructedDataAttribute struct {
	FcModelNode
}

func NewConstructedDataAttribute(objectReference *ObjectReference, fc string, children []*FcModelNode) *ConstructedDataAttribute {
	c := &ConstructedDataAttribute{}
	c.objectReference = objectReference
	c.Fc = fc
	c.children = make(map[string]*ModelNode)
	for _, child := range children {
		c.children[child.getName()] = (*ModelNode)(unsafe.Pointer(child))
		child.parent = (*ModelNode)(unsafe.Pointer(c))
	}

	return c
}
