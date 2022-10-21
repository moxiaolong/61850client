package main

import (
	"Go61850Client/src"
)

var (
	serverModel = src.ServerModel{}
)

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		log.Println(r)
	//	}
	//}()
	hostName := "localhost"
	port := 8080
	modelFilePath := "文件名"

	clientSap := src.NewClientSap()
	association := clientSap.Associate(hostName, port, src.NewEventListener())
	defer func() {
		association.Close()
	}()

	serverModel = src.SclParserParse(modelFilePath)[0]
	association.ServerModel = &serverModel

	serverModel = association.RetrieveModel()
}
