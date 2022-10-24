package src

type CalledPresentationSelector struct {
	value []byte
}

func NewCalledPresentationSelector(value []byte) *CalledPresentationSelector {
	return &CalledPresentationSelector{value: value}
}
