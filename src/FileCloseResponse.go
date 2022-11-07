package src

import "bytes"

type FileCloseResponse struct {
}

func (r FileCloseResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r FileCloseResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileCloseResponse() *FileCloseResponse {
	return &FileCloseResponse{}
}
