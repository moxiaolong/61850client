package src

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	bdaFloat32 := NewBdaFloat32(nil, "", "", false, false)
	bdaFloat32.SetFloat(3.3)
	//[8 64 83 51 51]
	fmt.Printf("%v", bdaFloat32.value)
}
