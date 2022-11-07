package src

import "bytes"

type FileOpenRequest struct {
}

func (r FileOpenRequest) decode(is *bytes.Buffer, b bool) int {

}

func (r FileOpenRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileOpenRequest() *FileOpenRequest {
	return &FileOpenRequest{}
}
