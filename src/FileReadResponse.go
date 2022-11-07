package src

import "bytes"

type FileReadResponse struct {
}

func (r FileReadResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r FileReadResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileReadResponse() *FileReadResponse {
	return &FileReadResponse{}
}
