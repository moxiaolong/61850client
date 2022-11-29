package src

import "fmt"

func HexStringFromBytes(bytes []byte) string {
	return fmt.Sprintf("%x", bytes)
}
