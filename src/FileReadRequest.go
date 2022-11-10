package src

type FileReadRequest struct {
	Integer32
}

func NewFileReadRequest() *FileReadRequest {
	return &FileReadRequest{}
}
