package src

type Result struct {
	BerInteger
}

func NewResult() *Result {
	return &Result{BerInteger: *NewBerInteger(nil, 0)}
}
