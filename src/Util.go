package src

import "bytes"

func readFully(is *bytes.Buffer, buffer []byte) {
	_, err := is.Read(buffer)
	if err != nil {
		panic(err)
	}

}
