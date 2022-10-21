package src

type AssociationInformation struct {
	Myexternal []*Myexternal
}

func NewAssociationInformation() *AssociationInformation {
	return &AssociationInformation{}
}
