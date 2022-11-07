package src

import "bytes"

type ObjectName struct {
}

func (n ObjectName) decode(is *bytes.Buffer, tag *BerTag) int {

}

func (n ObjectName) encode(os *ReverseByteArrayOutputStream) int {

}

func NewObjectName() *ObjectName {
	return &ObjectName{}
}
