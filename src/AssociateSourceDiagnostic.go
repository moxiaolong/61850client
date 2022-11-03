package src

import (
	"bytes"
)

type AssociateSourceDiagnostic struct {
}

func (d AssociateSourceDiagnostic) decode(is *bytes.Buffer, null interface{}) int {

}

func (d AssociateSourceDiagnostic) encode(os *ReverseByteArrayOutputStream) int {

}

func NewAssociateSourceDiagnostic() *AssociateSourceDiagnostic {
	return &AssociateSourceDiagnostic{}
}
