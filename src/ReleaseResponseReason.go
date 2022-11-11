package src

type ReleaseResponseReason struct {
	BerInteger
}

func NewReleaseResponseReason() *ReleaseResponseReason {
	return &ReleaseResponseReason{BerInteger: *NewBerInteger(nil, 0)}
}
