package main

import (
	"Go61850Client/src"
	"time"
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

	for {
		time.Sleep(time.Millisecond * 10)
	}
}

//func main() {
//
//	buffer := []byte{1, 2, 3, 4}
//	index := 0
//	subBuffer := buffer[index+1:]
//	fmt.Printf("%s", subBuffer)
//	//subBufferLength := len(buffer) - index - 1;
//
//	//byte[] subBuffer = new byte[subBufferLength];
//	//System.arraycopy(buffer, index + 1, subBuffer, 0, subBufferLength);
//	//System.out.println(Arrays.toString(subBuffer));
//
//}
