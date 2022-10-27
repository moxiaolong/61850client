package src

type AccessResult struct {
	Success *Data
}

func NewAccessResult() *AccessResult {
	return &AccessResult{}
}
