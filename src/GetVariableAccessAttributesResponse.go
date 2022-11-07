package src

import (
	"bytes"
)

type GetVariableAccessAttributesResponse struct {
	typeDescription *TypeDescription
}

func (r GetVariableAccessAttributesResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r GetVariableAccessAttributesResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewGetVariableAccessAttributesResponse() *GetVariableAccessAttributesResponse {
	return &GetVariableAccessAttributesResponse{}
}
