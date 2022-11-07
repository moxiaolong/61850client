package src

import "bytes"

type InformationReport struct {
	VariableAccessSpecification *VariableAccessSpecification
	listOfAccessResult          *ListOfAccessResult
}

func (r InformationReport) decode(is *bytes.Buffer, b bool) int {

}

func (r InformationReport) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewInformationReport() *InformationReport {
	return &InformationReport{}
}
