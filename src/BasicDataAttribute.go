package src

type BasicDataAttribute struct {
	FcModelNode
	basicType string
	sAddr     string
	dchg      bool
	dupd      bool
	qchg      bool
	chgRcbs   []*Urcb
	dupdRcbs  []*Urcb
}

func NewBasicDataAttribute(objectReference *ObjectReference, fc string, sAddr string, dchg bool, dupd bool) *BasicDataAttribute {
	b := &BasicDataAttribute{}
	b.FcModelNode = *NewFcModelNode()
	b.ObjectReference = objectReference
	b.Fc = fc
	b.sAddr = sAddr
	b.dchg = dchg
	b.dupd = dupd

	if dchg {
		b.chgRcbs = make([]*Urcb, 0)
	} else {
		b.chgRcbs = nil
	}
	if dupd {
		b.dupdRcbs = make([]*Urcb, 0)
	} else {
		b.dupdRcbs = nil
	}

	return b
}

func (a *BasicDataAttribute) getChild(childName string, fc string) *ModelNode {
	return nil
}
