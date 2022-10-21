package src

type AARQApdu struct {
	ApplicationContextName *BerObjectIdentifier
	CalledAPTitle          *APTitle
	CalledAEQualifier      *AEQualifier
	CallingAPTitle         *APTitle
	CallingAEQualifier     *AEQualifier
	UserInformation        *AssociationInformation
}

func NewAARQApdu() *AARQApdu {
	return &AARQApdu{}
}
