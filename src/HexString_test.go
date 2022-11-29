package src

import "testing"

func TestHexString(t *testing.T) {
	println(HexStringFromBytes([]byte{0, 0, 0, 1}))
}
