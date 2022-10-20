package src

var (
	serverModel = ServerModel{}
)

func main() {

	hostName := "localhost"
	port := 8080
	modelFilePath := "文件名"

	clientSap := newClientSap()
	association := clientSap.associate(hostName, port, NewEventListener())
	defer func() {
		association.close()
	}()

	serverModel = SclParserParse(modelFilePath)[0]
	association.ServerModel = &serverModel

	serverModel = association.retrieveModel()
}
