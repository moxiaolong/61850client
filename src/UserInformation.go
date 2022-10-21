package src

type UserInformation struct {
	Myexternal []*Myexternal
}

func NewUserInformation() *UserInformation {
	return &UserInformation{}
}
