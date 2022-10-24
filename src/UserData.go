package src

type UserData struct {
	FullyEncodedData *FullyEncodedData
}

func NewUserData() *UserData {
	return &UserData{}
}
