package src

type CallingPresentationSelector struct {
	PresentationSelector
}

func NewCallingPresentationSelector(value []byte) *CallingPresentationSelector {
	return &CallingPresentationSelector{PresentationSelector: *NewPresentationSelector(value)}
}
