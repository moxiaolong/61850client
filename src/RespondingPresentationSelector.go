package src

type RespondingPresentationSelector struct {
	PresentationSelector
}

func NewRespondingPresentationSelector(value []byte) *RespondingPresentationSelector {
	return &RespondingPresentationSelector{*NewPresentationSelector(value)}
}
