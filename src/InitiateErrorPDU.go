package src

type InitiateErrorPDU struct {
	ServiceError
}

func NewInitiateErrorPDU() *InitiateErrorPDU {
	serviceError := ()
	return &InitiateErrorPDU{ServiceError: *serviceError}
}
