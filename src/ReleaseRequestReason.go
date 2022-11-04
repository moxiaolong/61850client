package src

type ReleaseRequestReason struct {
	BerInteger
}

func NewReleaseRequestReason() *ReleaseRequestReason {
	return &ReleaseRequestReason{}
}
