package src

type CPTypeNormalModeParameters struct {
	CallingPresentationSelector       *CallingPresentationSelector
	CalledPresentationSelector        *CalledPresentationSelector
	PresentationContextDefinitionList *PresentationContextDefinitionList
	UserData                          *UserData
}

func NewCPTypeNormalModeParameters() *CPTypeNormalModeParameters {
	return &CPTypeNormalModeParameters{}
}
