package src

type FileCloseRequest struct {
	Integer32
}

func NewFileCloseRequest() *FileCloseRequest {
	return &FileCloseRequest{Integer32: *NewInteger32(0)}
}
