package src

type ClientTSap struct {
	acseSap                *ClientAcseSap
	tSelLocal              []byte
	tSelRemote             []byte
	MaxTPDUSizeParam       int
	MessageFragmentTimeout int
	MessageTimeout         int
}

func NewClientTSap() *ClientTSap {
	return &ClientTSap{}
}
