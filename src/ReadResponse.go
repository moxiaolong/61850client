package src

import "bytes"

type ReadResponse struct {
}

func (r ReadResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r ReadResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewReadResponse() *ReadResponse {
	return &ReadResponse{}
}
