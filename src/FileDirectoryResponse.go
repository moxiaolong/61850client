package src

import "bytes"

type FileDirectoryResponse struct {
}

func (r FileDirectoryResponse) decode(is *bytes.Buffer, b bool) int {

}

func (r FileDirectoryResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileDirectoryResponse() *FileDirectoryResponse {
	return &FileDirectoryResponse{}
}
