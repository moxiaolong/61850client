package src

type ApplicationContextName struct {
	BerObjectIdentifier
}

func NewApplicationContextName() *ApplicationContextName {
	return &ApplicationContextName{BerObjectIdentifier: *NewBerObjectIdentifier(nil)}
}
