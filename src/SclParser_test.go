package src

import (
	"io/ioutil"
	"testing"
)

func TestParseStream(t *testing.T) {
	parser := SclParser{}
	file, err := ioutil.ReadFile("C:\\Users\\DragonMo\\GolandProjects\\Go61850Client\\iec61850bean-sample01.icd")
	if err != nil {
		panic(err)
	}
	err = parser.ParseStream(file)
}
