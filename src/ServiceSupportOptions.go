package src

type ServiceSupportOptions struct {
	BerBitString
}

func NewServiceSupportOptions(value []byte, numBits int) *ServiceSupportOptions {
	return &ServiceSupportOptions{BerBitString: *NewBerBitString(nil, value, numBits)}
}
