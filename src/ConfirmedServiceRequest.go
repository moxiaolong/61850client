package src

import "bytes"

type ConfirmedServiceRequest struct {
}

func (r ConfirmedServiceRequest) decode(is *bytes.Buffer, tag *BerTag) int {

}

func (r ConfirmedServiceRequest) encode(os *ReverseByteArrayOutputStream) int {

}

func NewConfirmedServiceRequest() *ConfirmedServiceRequest {
	return &ConfirmedServiceRequest{}
}
