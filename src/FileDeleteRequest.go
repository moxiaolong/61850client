package src

type FileDeleteRequest struct {
	FileName
}

func NewFileDeleteRequest() *FileDeleteRequest {
	return &FileDeleteRequest{FileName: *NewFileName()}
}
