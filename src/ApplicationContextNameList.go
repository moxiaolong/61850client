package src

import "bytes"

type ApplicationContextNameList struct {
}

func (l ApplicationContextNameList) decode(is *bytes.Buffer, b bool) int {

}

func (l ApplicationContextNameList) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewApplicationContextNameList() *ApplicationContextNameList {
	return &ApplicationContextNameList{}
}
