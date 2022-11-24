package src

import "strings"

type DataSet struct {
	Members          []*FcModelNode
	DataSetReference string
	deletable        bool
	// Map<Fc, Map<String, FcModelNode>> MembersMap
	MembersMap map[string]map[string]*FcModelNode
}

func (s *DataSet) getMembers() []*FcModelNode {
	return s.Members
}

func NewDataSet() *DataSet {
	return &DataSet{}
}
func NewDataSetWithRef(dataSetReference string, members []*FcModelNode, deletable bool) *DataSet {
	d := &DataSet{}
	if !strings.HasPrefix(dataSetReference, "@") && strings.Index(dataSetReference, "/") == -1 {
		throw(
			"DataSet reference " + dataSetReference + " is invalid. Must either start with @ or contain a slash.")
	}
	d.Members = make([]*FcModelNode, 0)
	d.DataSetReference = dataSetReference
	d.deletable = deletable
	d.MembersMap = make(map[string]map[string]*FcModelNode)

	//TODO

	for _, member := range members {
		d.Members = append(d.Members, member)
		if d.MembersMap[member.Fc] == nil {
			d.MembersMap[member.Fc] = make(map[string]*FcModelNode)
		}
		d.MembersMap[member.Fc][member.ObjectReference.toString()] = member
	}

	return d
}
