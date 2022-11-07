package src

type FileCloseResponse struct {
	BerNull
}

func NewFileCloseResponse() *FileCloseResponse {
	return &FileCloseResponse{}
}
