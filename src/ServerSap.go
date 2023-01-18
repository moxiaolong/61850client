package src

type ServerSap struct {
	serverModel *ServerModel
	port        int
	backlog     int
}

func NewServerSap(port int,
	backlog int,
	serverModel *ServerModel) *ServerSap {
	return &ServerSap{port: port, backlog: backlog, serverModel: serverModel}
}
