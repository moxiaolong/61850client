package src

type CalledPresentationSelector struct {
	PresentationSelector
}

func NewCalledPresentationSelector(value []byte) *CalledPresentationSelector {
	return &CalledPresentationSelector{PresentationSelector: *NewPresentationSelector(value)}
}
