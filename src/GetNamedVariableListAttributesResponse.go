package src

import (
	"bytes"
)

type GetNamedVariableListAttributesResponse struct {
}

func (r GetNamedVariableListAttributesResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r GetNamedVariableListAttributesResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewGetNamedVariableListAttributesResponse() *GetNamedVariableListAttributesResponse {
	return &GetNamedVariableListAttributesResponse{}
}
