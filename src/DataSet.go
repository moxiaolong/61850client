package src

type DataSet struct {
}

func (s *DataSet) getMembers() []*FcModelNode {
	return nil
}

func NewDataSet() *DataSet {
	return &DataSet{}
}
