package src

type UnconfirmedService struct {
	InformationReport *InformationReport
}

func NewUnconfirmedService() *UnconfirmedService {
	return &UnconfirmedService{}
}
