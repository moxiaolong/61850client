package src

type ServerThread struct {
}

func (t ServerThread) connectionClosedSignal() {

}

func NewServerThread() *ServerThread {
	return &ServerThread{}
}
