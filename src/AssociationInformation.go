package src

type AssociationInformation struct {
	Myexternal []*Myexternal
}

func (i AssociationInformation) encode(os *ReverseByteArrayOutputStream, b bool) int {
	//TODO
	return 0
}

func NewAssociationInformation() *AssociationInformation {
	return &AssociationInformation{}
}
