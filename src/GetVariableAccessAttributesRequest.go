package src

import "bytes"

type GetVariableAccessAttributesRequest struct {
}

func (r GetVariableAccessAttributesRequest) decode(is *bytes.Buffer, t interface{}) int {

}

func (r GetVariableAccessAttributesRequest) encode(os *ReverseByteArrayOutputStream) int {

}

func NewGetVariableAccessAttributesRequest() *GetVariableAccessAttributesRequest {
	return &GetVariableAccessAttributesRequest{}
}
