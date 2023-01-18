package src

import "github.com/beevik/etree"

type Da struct {
	AbstractDataAttribute
	fc   string
	dchg bool
	dupd bool
	qchg bool
}

func (d *Da) getFc() string {
	return d.fc
}

func (d *Da) isDchg() bool {
	return d.dchg
}

func (d *Da) isDupd() bool {
	return d.dupd
}

func (d *Da) isQchg() bool {
	return d.qchg
}

func NewDa(node *etree.Element) *Da {
	attribute := NewAbstractDataAttribute(node)
	d := &Da{AbstractDataAttribute: *attribute}
	for _, node := range node.Attr {
		nodeName := node.Key

		if nodeName == ("fc") {
			d.fc = node.Value
		} else if nodeName == ("dchg") {
			d.dchg = "true" == (node.Value)
		} else if nodeName == ("qchg") {
			d.qchg = "true" == (node.Value)
		} else if nodeName == ("dupd") {
			d.dupd = "true" == (node.Value)
		}
	}
	return d
}
