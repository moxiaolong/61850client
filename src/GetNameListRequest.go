package src

import (
	"bytes"
)

type GetNameListRequest struct {
}

func (r GetNameListRequest) decode(is *bytes.Buffer, b bool) int {

}

func (r GetNameListRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewGetNameListRequest() *GetNameListRequest {
	return &GetNameListRequest{}
}
