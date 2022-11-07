package src

import (
	"bytes"
)

type GetNameListResponse struct {
}

func (r GetNameListResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r GetNameListResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewGetNameListResponse() *GetNameListResponse {
	return &GetNameListResponse{}
}
