package src

type ParameterSupportOptions struct {
	BerBitString
}

func NewParameterSupportOptions(code []byte) *ParameterSupportOptions {
	bitString := NewBerBitString(code, nil, 0)
	return &ParameterSupportOptions{BerBitString: *bitString}
}
