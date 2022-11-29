package src

import "strconv"

type ConstructedDataAttribute struct {
	FcModelNode
}

func (c *ConstructedDataAttribute) setValueFromMmsDataObj(data *Data) {
	if data.structure == nil {
		throw("ServiceError.TYPE_CONFLICT expected type: structure")
	}
	if len(data.structure.seqOf) != len(c.Children) {
		throw("ServiceError.TYPE_CONFLICT expected type: structure with " + strconv.Itoa(len(c.Children)) + " elements")
	}

	i := 0
	for _, child := range c.Children {
		child.setValueFromMmsDataObj(data.structure.seqOf[i])
		i++
	}

}

func (c *ConstructedDataAttribute) copy() ModelNodeI {
	subDataAttributesCopy := make([]ModelNodeI, 0)
	for _, subDA := range c.Children {
		subDataAttributesCopy = append(subDataAttributesCopy, subDA.copy())
	}
	return NewConstructedDataAttribute(c.getObjectReference(), c.Fc, subDataAttributesCopy)
}

func NewConstructedDataAttribute(objectReference *ObjectReference, fc string, children []ModelNodeI) *ConstructedDataAttribute {
	c := &ConstructedDataAttribute{}
	c.ObjectReference = objectReference
	c.Fc = fc
	c.Children = make(map[string]ModelNodeI)
	for _, child := range children {
		c.Children[child.getName()] = child
		child.setParent(c)
	}

	return c
}
