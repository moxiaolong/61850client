package src

import "bytes"

type FileReadRequest struct {
}

func (r FileReadRequest) decode(is *bytes.Buffer, b bool) int {

}

func (r FileReadRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileReadRequest() *FileReadRequest {
	return &FileReadRequest{}
}
