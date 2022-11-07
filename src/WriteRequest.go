package src

import "bytes"

type WriteRequest struct {
}

func (r WriteRequest) decode(is *bytes.Buffer, b bool) int {

}

func (r WriteRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewWriteRequest() *WriteRequest {
	return &WriteRequest{}
}
