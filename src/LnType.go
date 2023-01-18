package src

import "github.com/beevik/etree"

type LnType struct {
	AbstractType
	dos []*Do
}

func NewLnType(node *etree.Element) *LnType {
	abstractType := NewAbstractType(node)
	l := &LnType{AbstractType: *abstractType}
	if node.SelectAttr("lnClass") == nil {
		panic("Required attribute \"lnClass\" not found in LNType!")
	}

	for _, element := range node.ChildElements() {
		if element.Tag == ("DO") {
			l.dos = append(l.dos, NewDo(element))
		}
	}

	return l

}
