package src

type AAREApdu struct {
	UserInformation *UserInformation
}

func NewAAREApdu() *AAREApdu {
	return &AAREApdu{}
}
