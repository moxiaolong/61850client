package src

type InitiateErrorPDU struct {
	ServiceError
}

func NewInitiateErrorPDU() *InitiateErrorPDU {
	serviceError := NewServiceError()
	return &InitiateErrorPDU{ServiceError: *serviceError}

}
