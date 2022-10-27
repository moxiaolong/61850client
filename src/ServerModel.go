package src

type ServerModel struct {
	ModelNode
	urcbs map[string]*Urcb
}

func SclParserParse(path string) []ServerModel {
	return make([]ServerModel, 1)
}
