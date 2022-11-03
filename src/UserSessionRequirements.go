package src

type UserSessionRequirements struct {
	BerBitString
}

func NewUserSessionRequirements() *UserSessionRequirements {
	return &UserSessionRequirements{}
}
