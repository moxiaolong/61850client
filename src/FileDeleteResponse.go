package src

import "bytes"

type FileDeleteResponse struct {
}

func (r FileDeleteResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r FileDeleteResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileDeleteResponse() *FileDeleteResponse {
	return &FileDeleteResponse{}
}
