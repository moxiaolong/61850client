package src

import "bytes"

type FileAttributes struct {
}

func (a FileAttributes) decode(is *bytes.Buffer, b bool) int {

}

func (a FileAttributes) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewFileAttributes() *FileAttributes {
	return &FileAttributes{}
}
