package src

type DataAccessError struct {
	BerInteger
}

func NewDataAccessError() *DataAccessError {
	return &DataAccessError{}
}
