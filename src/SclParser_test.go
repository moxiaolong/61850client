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
			if fcModelNode.getChildren() != nil && len(fcModelNode.getChildren()) > 0 {
				for _, item := range fcModelNode.getChildren() {
					fcNodeMap := make(map[string]string)
					fcNodeMap["ref"] = fc + "$" + item.getObjectReference().toString()
					fcNodeMap["desc"] = item.getDesc()
					result[referenceStr] = append(result[referenceStr], fcNodeMap)

				}
			} else {
				fcNodeMap := make(map[string]string)
				fcNodeMap["ref"] = fc + "$" + name
				fcNodeMap["desc"] = fcModelNode.getDesc()
				result[referenceStr] = append(result[referenceStr], fcNodeMap)
			}
		}

	}

	return result
}
