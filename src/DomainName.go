package src

import "bytes"

type DomainName struct {
}

func (n DomainName) decode(is *bytes.Buffer, t interface{}) int {

}

func (n DomainName) encode(os *ReverseByteArrayOutputStream) int {

}

func NewDomainName() *DomainName {
	return &DomainName{}
}
