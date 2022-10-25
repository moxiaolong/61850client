package src

type InitiateErrorPDU struct {
	ServiceError
	additionalDescription *BerVisibleString
	additionalCode        *BerInteger
}

func NewInitiateErrorPDU() *InitiateErrorPDU {
	serviceError := NewServiceError()
	return &InitiateErrorPDU{ServiceError: *serviceError}

}
