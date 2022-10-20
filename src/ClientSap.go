package src

type ClientSap struct {
	acseSap *ClientAcseSap
}

func newClientSap() *ClientSap {
	r := &ClientSap{}
	acseSap := newClientAcseSap()
	acseSap.tSap.tSelLocal = []byte{0, 0}
	acseSap.tSap.tSelRemote = []byte{0, 1}
	acseSap.tSap.MaxTPDUSizeParam = 10
	r.acseSap = acseSap
	return r
}

func (c *ClientSap) associate(address string, port int, eventListener *EventListener) *ClientAssociation {
	clientAssociation :=
		NewClientAssociation(
			address,
			port,
			c.acseSap,
			65000,
			5,
			5,
			10,
			[]byte{0xee, 0x1c, 0, 0, 0x04, 0x08, 0, 0, 0x79, 0xef, 0x18},
			20000,
			10000,
			eventListener)
	return clientAssociation
}
