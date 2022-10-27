package src

type InformationReport struct {
	VariableAccessSpecification *VariableAccessSpecification
	listOfAccessResult          *ListOfAccessResult
}

func NewInformationReport() *InformationReport {
	return &InformationReport{}
}
