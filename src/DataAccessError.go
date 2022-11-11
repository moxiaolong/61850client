package src

type DataAccessError struct {
	BerInteger
}

func NewDataAccessError() *DataAccessError {
	return &DataAccessError{BerInteger: *NewBerInteger(nil, 0)}
}
