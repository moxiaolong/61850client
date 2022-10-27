package src

type PresentationContextDefinitionList struct {
	ContextList
}

func NewPresentationContextDefinitionList(code []byte) *PresentationContextDefinitionList {
	return &PresentationContextDefinitionList{ContextList: *NewContextList(code)}
}
