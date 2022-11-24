package src

import "unsafe"

type FcDataObject struct {
	FcModelNode
}

func NewFcDataObject(objectReference *ObjectReference, fc string, children []*FcModelNode) *FcDataObject {
	f := &FcDataObject{}
	f.Children = make(map[string]*ModelNode)
	f.ObjectReference = objectReference
	for _, child := range children {
		f.Children[child.ObjectReference.getName()] = (*ModelNode)(unsafe.Pointer(child))
		child.parent = (*ModelNode)(unsafe.Pointer(f))
	}
	f.Fc = fc

	return f
}
