package src

type VariableAccessSpecification struct {
	ListOfVariable *VariableDefs
}

func NewVariableAccessSpecification() *VariableAccessSpecification {
	return &VariableAccessSpecification{}
}
