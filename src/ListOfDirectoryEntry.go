package src

import "bytes"

type ListOfDirectoryEntry struct {
}

func (e ListOfDirectoryEntry) decode(is *bytes.Buffer, b bool) int {

}

func (e ListOfDirectoryEntry) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewListOfDirectoryEntry() *ListOfDirectoryEntry {
	return &ListOfDirectoryEntry{}
}
