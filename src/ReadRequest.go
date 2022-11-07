package src

import "bytes"

type ReadRequest struct {
}

func (r ReadRequest) decode(is *bytes.Buffer, b bool) int {

}

func (r ReadRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewReadRequest() *ReadRequest {
	return &ReadRequest{}
}
