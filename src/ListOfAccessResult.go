package src

type ListOfAccessResult struct {
	seqOf []*AccessResult
}

func NewListOfAccessResult() *ListOfAccessResult {
	return &ListOfAccessResult{}
}
