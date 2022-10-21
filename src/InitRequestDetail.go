package src

type InitRequestDetail struct {
	ServicesSupportedCalling *ServiceSupportOptions
	ProposedParameterCBB     *ParameterSupportOptions
	ProposedVersionNumber    *Integer16
}

func NewInitRequestDetail() *InitRequestDetail {
	return &InitRequestDetail{}
}
