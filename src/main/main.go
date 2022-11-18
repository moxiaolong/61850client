package main

import (
	"Go61850Client/src"
	"time"
)

var (
	serverModel = &src.ServerModel{}
)

func main() {

	//defer func() {
	//	if r := recover(); r != nil {
	//		log.Println(r)
	//	}
	//}()
	hostName := "localhost"
	port := 8080
	//modelFilePath := "文件名"

	clientSap := src.NewClientSap()
	association := clientSap.Associate(hostName, port, src.NewEventListener())
	defer func() {
		association.Close()
	}()

	//serverModel = src.SclParserParse(modelFilePath)[0]
	//association.ServerModel = serverModel

	serverModel = association.RetrieveModel()

	for {
		time.Sleep(time.Millisecond * 10)
	}
}

//func main() {
//
//	buffer := []byte{0, 0, 0, 0, 0, 0, 0, 5}
//	index := 6
//	newBuffer := make([]byte, len(buffer)*2)
//	for i, b := range buffer {
//		newBuffer[len(buffer)+i] = b
//	}
//	index += len(buffer)
//	buffer = newBuffer
//
//	fmt.Printf("%x\r\n", buffer)
//	fmt.Printf("%d\r\n", index)
//
//	//byteArray := []byte{1, 2, 3, 4}
//	//buffer := []byte{0, 0, 0, 0, 0, 0, 0, 5}
//	//index := 6
//	//
//	//for i := len(byteArray) - 1; i >= 0; i-- {
//	//	buffer[index] = byteArray[i]
//	//	index -= 1
//	//}
//	//
//	//fmt.Printf("%x", buffer)
//	//subBufferLength := len(buffer) - index - 1;
//
//	//byte[] subBuffer = Newbyte[subBufferLength];
//	//System.arraycopy(buffer, index + 1, subBuffer, 0, subBufferLength);
//	//System.out.println(Arrays.toString(subBuffer));
//
//}
