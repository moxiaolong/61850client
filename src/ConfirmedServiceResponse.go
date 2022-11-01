package src

type ConfirmedServiceResponse struct {
	GetVariableAccessAttributesResponse *GetVariableAccessAttributesResponse
	GetVariableAccessAttributes         *GetVariableAccessAttributes
}

func NewConfirmedServiceResponse() *ConfirmedServiceResponse {
	return &ConfirmedServiceResponse{}
}
