package src

import "bytes"

type FileDirectoryRequest struct {
}

func (r FileDirectoryRequest) decode(is *bytes.Buffer, b bool) int {

}

func (r FileDirectoryRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileDirectoryRequest() *FileDirectoryRequest {
	return &FileDirectoryRequest{}
}
