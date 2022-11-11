package src

type GetNamedVariableListAttributesRequest struct {
	ObjectName
}

func NewGetNamedVariableListAttributesRequest() *GetNamedVariableListAttributesRequest {
	return &GetNamedVariableListAttributesRequest{ObjectName: *NewObjectName()}
}
