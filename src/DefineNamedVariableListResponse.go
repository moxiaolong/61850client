package src

type DefineNamedVariableListResponse struct {
	BerNull
}

func NewDefineNamedVariableListResponse() *DefineNamedVariableListResponse {
	return &DefineNamedVariableListResponse{BerNull: *NewBerNull()}
}
