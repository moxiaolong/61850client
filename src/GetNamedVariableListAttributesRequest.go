package src

import "bytes"

type GetNamedVariableListAttributesRequest struct {
}

func (r GetNamedVariableListAttributesRequest) decode(is *bytes.Buffer, t interface{}) int {

}

func (r GetNamedVariableListAttributesRequest) encode(os *ReverseByteArrayOutputStream) int {

}

func NewGetNamedVariableListAttributesRequest() *GetNamedVariableListAttributesRequest {
	return &GetNamedVariableListAttributesRequest{}
}
