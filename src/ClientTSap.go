package src

type ClientTSap struct {
	acseSap                *ClientAcseSap
	tSelLocal              []byte
	tSelRemote             []byte
	MaxTPDUSizeParam       int
	MessageFragmentTimeout int
	MessageTimeout         int
}

func (s *ClientTSap) connectTo(address string, port int) *TConnection {
	return &TConnection{}
}

func NewClientTSap() *ClientTSap {
	return &ClientTSap{}
}
