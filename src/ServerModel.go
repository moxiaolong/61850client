package src

type ServerModel struct {
	ModelNode
	urcbs map[string]*Urcb
	brcbs map[string]*Brcb
}

func (m ServerModel) getDataSet(ref string) *DataSet {

}

func SclParserParse(path string) []ServerModel {
	return make([]ServerModel, 1)
}
