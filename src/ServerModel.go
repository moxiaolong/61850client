package src

type ServerModel struct {
	ModelNode
	urcbs map[string]*Urcb
	brcbs map[string]*Brcb
}

func NewServerModel([]*LogicalDevice, []*DataSet) *ServerModel {
	return &ServerModel{}

}

func (m ServerModel) getDataSet(ref string) *DataSet {
	return nil
}

func SclParserParse(path string) []*ServerModel {
	return make([]*ServerModel, 1)
}
