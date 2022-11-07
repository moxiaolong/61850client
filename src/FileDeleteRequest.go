package src

import "bytes"

type FileDeleteRequest struct {
}

func (r FileDeleteRequest) decode(is *bytes.Buffer, b bool) int {

}

func (r FileDeleteRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileDeleteRequest() *FileDeleteRequest {
	return &FileDeleteRequest{}
}
