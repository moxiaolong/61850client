package src

import "bytes"

type FileOpenResponse struct {
}

func (r FileOpenResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r FileOpenResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileOpenResponse() *FileOpenResponse {
	return &FileOpenResponse{}
}
