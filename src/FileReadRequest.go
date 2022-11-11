package src

type FileReadRequest struct {
	Integer32
}

func NewFileReadRequest() *FileReadRequest {
	return &FileReadRequest{Integer32: *NewInteger32(0)}
}
