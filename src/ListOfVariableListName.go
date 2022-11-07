package src

import "bytes"

type ListOfVariableListName struct {
}

func (n ListOfVariableListName) decode(is *bytes.Buffer, b bool) int {

}

func (n ListOfVariableListName) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewListOfVariableListName() *ListOfVariableListName {
	return &ListOfVariableListName{}
}
