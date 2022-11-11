package src

type ReleaseRequestReason struct {
	BerInteger
}

func NewReleaseRequestReason() *ReleaseRequestReason {
	return &ReleaseRequestReason{BerInteger: *NewBerInteger(nil, 0)}
}
