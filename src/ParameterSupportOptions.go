package src

type ParameterSupportOptions struct {
}

func (o *ParameterSupportOptions) encode(os *ReverseByteArrayOutputStream, b bool) int {
	return 0

}

func NewParameterSupportOptions([]byte) *ParameterSupportOptions {
	return &ParameterSupportOptions{}
}
