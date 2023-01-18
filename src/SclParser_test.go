package src

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseStream(t *testing.T) {
	parser := NewSclParser()
	file, err := ioutil.ReadFile("C:\\Users\\DragonMo\\GolandProjects\\Go61850Client\\iec61850bean-sample01.icd")
	if err != nil {
		panic(err)
	}
	err = parser.ParseStream(file)
	if err != nil {
		panic(err)
	}
	model := parser.ServerModel
	toMap := serverModelToMap(model)
	marshal, err := json.Marshal(toMap)
	println(string(marshal))
}

func serverModelToMap(serverModel *ServerModel) map[string][]map[string]string {
	result := make(map[string][]map[string]string)
	for _, dataSet := range serverModel.DataSets {
		referenceStr := dataSet.DataSetReference
		nameList := make([]map[string]string, 0)
		result[referenceStr] = nameList
		for _, fcModelNode := range dataSet.Members {
			fc := fcModelNode.getFc()
			name := fcModelNode.getObjectReference().toString()
			fcNodeMap := make(map[string]string)
			fcNodeMap["ref"] = fc + "$" + name
			fcNodeMap["desc"] = fcModelNode.getDesc()
			result[referenceStr] = append(result[referenceStr], fcNodeMap)
		}

	}

	return result
}
