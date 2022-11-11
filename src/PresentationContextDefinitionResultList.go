package src

type PresentationContextDefinitionResultList struct {
	ResultList
}

func NewPresentationContextDefinitionResultList() *PresentationContextDefinitionResultList {
	return &PresentationContextDefinitionResultList{ResultList: *NewResultList()}
}
