package src

type MechanismName struct {
	BerObjectIdentifier
}

func NewMechanismName() *MechanismName {
	return &MechanismName{}
}
