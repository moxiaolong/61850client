package src

import (
	"fmt"
	"net"
)

type ClientTSap struct {
	acseSap                *ClientAcseSap
	tSelLocal              []byte
	tSelRemote             []byte
	MaxTPDUSizeParam       int
	MessageFragmentTimeout int
	MessageTimeout         int
	serverThread           *ServerThread
}

func (c *ClientTSap) connectTo(address string, port int) *TConnection {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		panic(err)
	}

	tConnection := NewTConnection(&conn, c.MaxTPDUSizeParam, c.MessageTimeout, c.MessageFragmentTimeout, nil)
	tConnection.TSelRemote = c.tSelRemote
	tConnection.TSelLocal = c.tSelLocal
	tConnection.startConnection()
	return tConnection
}

func NewClientTSap() *ClientTSap {
	return &ClientTSap{}
}
