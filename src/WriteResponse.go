package src

import "bytes"

type WriteResponse struct {
}

func (r WriteResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r WriteResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewWriteResponse() *WriteResponse {
	return &WriteResponse{}
}
